#!/bin/sh

# # Prerequisites
# * `swig`
# * `patch`
# * `wget`
# * `unzip`
# * `sha256sum`

set -e

OUT_DIR=$1

if [ -z "${OUT_DIR}" ]; then
	echo Usage: generate.sh [out_dir]
	exit 1
fi

mkdir -p ${OUT_DIR}

# Cleanup old generated files
rm -rf \
	${OUT_DIR}/sampgdk.c \
	${OUT_DIR}/sampgdk.h \
	${OUT_DIR}/bishop.go \
	${OUT_DIR}/sampgdk_wrap.c \
	${OUT_DIR}/callbacks.go \
	${OUT_DIR}/sclinux.go \
	${OUT_DIR}/getch.go \
	${OUT_DIR}/amx

# Download and unarchive samp-plugin-sdk
SAMP_PLUGIN_SDK_VERSION=938781bafd8089cdb26af0617af24f862b2c3831
SAMP_PLUGIN_SDK_CHECKSUM=505cb49ed8af8c9dffe5cc35385276cd502316911c3b8b373a51cac2960f2f20
wget https://github.com/Zeex/samp-plugin-sdk/archive/${SAMP_PLUGIN_SDK_VERSION}.zip -O ${OUT_DIR}/samp-plugin-sdk.zip; \
	echo "${SAMP_PLUGIN_SDK_CHECKSUM} ${OUT_DIR}/samp-plugin-sdk.zip" | sha256sum -c -; \
	unzip -d ${OUT_DIR} ${OUT_DIR}/samp-plugin-sdk.zip; \
	rm ${OUT_DIR}/samp-plugin-sdk.zip; \
	mv ${OUT_DIR}/samp-plugin-sdk-${SAMP_PLUGIN_SDK_VERSION} ${OUT_DIR}/samp-plugin-sdk

# Download sampgdk amalgamation
SAMPGDK_VERSION=4.6.2
SAMPGDK_CHECKSUM=edebd36b03e5c809ac8f5992147988121bd7c8a0860207acee76db7e029d899f
wget https://github.com/Zeex/sampgdk/releases/download/v${SAMPGDK_VERSION}/sampgdk-${SAMPGDK_VERSION}-amalgamation.zip -O ${OUT_DIR}/sampgdk.zip; \
	echo "${SAMPGDK_CHECKSUM} ${OUT_DIR}/sampgdk.zip" | sha256sum -c -; \
	unzip -d ${OUT_DIR} ${OUT_DIR}/sampgdk.zip; \
	rm ${OUT_DIR}/sampgdk.zip; \
	rm ${OUT_DIR}/how-to-use-amalgamation.txt

# Get script dir
# Line below was shamelessly stolen from https://stackoverflow.com/a/246128
SRC_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Patch the header file for CGO
patch ${OUT_DIR}/sampgdk.h ${SRC_DIR}/sampgdk.h.patch.macros
patch ${OUT_DIR}/sampgdk.h ${SRC_DIR}/sampgdk.h.patch.callbacksremoval
patch ${OUT_DIR}/sampgdk.h ${SRC_DIR}/sampgdk.h.patch.exportsremoval
patch ${OUT_DIR}/sampgdk.h ${SRC_DIR}/sampgdk.h.patch.amxinternal
patch ${OUT_DIR}/sampgdk.c ${SRC_DIR}/sampgdk.c.patch.amxinternal

# Copy necessary headers from samp-plugin-sdk to out dir
cp ${OUT_DIR}/samp-plugin-sdk/plugincommon.h ${OUT_DIR}/
mkdir -p ${OUT_DIR}/amx
cp ${OUT_DIR}/samp-plugin-sdk/amx/amx.h ${OUT_DIR}/amx/
# Copy other headers from amx to the root dir itself because go build searches only in root folder
# by default (otherwise user would have to import them with -I during buildtime)
cp ${OUT_DIR}/samp-plugin-sdk/amx/sclinux.h ${OUT_DIR}/
cp ${OUT_DIR}/samp-plugin-sdk/amx/getch.h ${OUT_DIR}/

# Generate Go bindings with SWIG
swig -go -cgo -intgosize 32 -I${OUT_DIR} -outdir ${OUT_DIR} -o ${OUT_DIR}/sampgdk_wrap.c ${SRC_DIR}/bishop.i

# Remove sampgdk prefix from all functions
sed -i 's/func Sampgdk_/func /g' ${OUT_DIR}/bishop.go

# generate callbacks.go
CALLBACKS_FILENAME=${OUT_DIR}/callbacks.go

cat ${SRC_DIR}/callbacks.go.part > ${CALLBACKS_FILENAME}
LINES=$(cat ${SRC_DIR}/sampgdk.h.patch.callbacksremoval | grep -Po 'SAMPGDK_CALLBACK\(.*\)')
IFS=$'\n'
for LINE in ${LINES}; do
	LINE=$(echo ${LINE} | awk -F, '{for(i=2; i<=NF; i++) {print $i} print $1}')
	FN_NAME=$(echo ${LINE} | awk -F\( -v N=1 {'print $N'} | xargs)
	CGO_FN_DECL="func ${FN_NAME}("
	CGO_FN_BODY="Callback${FN_NAME}("
	GO_FN_DECL="func("
	GO_FN_BODY=""
	COUNTER=0
	IFS=$'\n'
	for PART in ${LINE}; do
		if [ ${COUNTER} -eq 0 ]; then
			PART=$(echo ${PART} | xargs | sed -e "s/^$FN_NAME(//")
		fi
		PART=$(echo ${PART} | xargs | sed -e "s/))$//")

		if [[ "${PART}" == SAMPGDK_CALLBACK* ]]; then
			RET_TYPE=$(echo ${PART} | sed -e "s/^SAMPGDK_CALLBACK(//")
			CGO_FN_DECL=$(echo ${CGO_FN_DECL} | sed -e 's/.\{2\}$//')
			GO_FN_DECL=$(echo ${GO_FN_DECL} | sed -e 's/.\{2\}$//')
			CGO_FN_BODY=$(echo ${CGO_FN_BODY} | sed -e 's/.\{2\}$//')
			if [ "${RET_TYPE}" = "int" ]; then
				GO_FN_BODY="return 1"
			elif [ "${RET_TYPE}" = "float" ]; then
				GO_FN_BODY="return 1"
			elif [ "${RET_TYPE}" = "bool" ]; then
				GO_FN_BODY="return true"
			fi

			if [ "${RET_TYPE}" = "void" ]; then
				CGO_FN_DECL=$(echo "${CGO_FN_DECL})")
				GO_FN_DECL=$(echo "${GO_FN_DECL})")
				CGO_FN_BODY=$(echo "${CGO_FN_BODY})")
			else
				CGO_FN_DECL=$(echo "${CGO_FN_DECL}) ${RET_TYPE}")
				GO_FN_DECL=$(echo "${GO_FN_DECL}) ${RET_TYPE}")
				CGO_FN_BODY=$(echo "return ${CGO_FN_BODY})")
			fi
		else
			GO_ARG_TYPE=$(echo ${PART} | sed -e '$s/\w*$//' | xargs)
			ARG_NAME=$(echo ${PART} | sed -e "s/^$GO_ARG_TYPE//" | sed -e "s/^\*//" | xargs)
			if [ "${ARG_NAME}" = "type" ]; then
				ARG_NAME=t
			fi

			VAL=${ARG_NAME}
			if [ "${GO_ARG_TYPE}" = "float" ]; then
				GO_ARG_TYPE=float32
				CGO_ARG_TYPE=C.float
				VAL=$(echo "float32(${VAL})")
			elif [ "${GO_ARG_TYPE}" = "int" ]; then
				GO_ARG_TYPE=int
				CGO_ARG_TYPE=C.int
				VAL=$(echo "int(${VAL})")
			elif [ "${GO_ARG_TYPE}" = "const char *" ]; then
				GO_ARG_TYPE=string
				CGO_ARG_TYPE=*C.char
				VAL=$(echo "C.GoString(${VAL})")
			elif [ "${GO_ARG_TYPE}" = "bool" ]; then
				VAL=$(echo "int(${VAL}) != 0")
				CGO_ARG_TYPE=C.int
				GO_ARG_TYPE=bool
			fi
			CGO_FN_DECL=$(echo "${CGO_FN_DECL}${ARG_NAME} ${CGO_ARG_TYPE}, ")
			GO_FN_DECL=$(echo "${GO_FN_DECL}${ARG_NAME} ${GO_ARG_TYPE}, ")
			CGO_FN_BODY=$(echo "${CGO_FN_BODY}${VAL}, ")
		fi

		COUNTER=$((COUNTER+1))
	done

	echo var Callback${FN_NAME} = ${GO_FN_DECL} { ${GO_FN_BODY} } >> ${CALLBACKS_FILENAME}
	echo >> ${CALLBACKS_FILENAME}
	echo //export ${FN_NAME} >> ${CALLBACKS_FILENAME}
	echo ${CGO_FN_DECL} { >> ${CALLBACKS_FILENAME}
	echo "	${CGO_FN_BODY}" >> ${CALLBACKS_FILENAME}
	echo } >> ${CALLBACKS_FILENAME}
	echo >> ${CALLBACKS_FILENAME}
done

# # Remove samp-plugin-sdk
rm -rf ${OUT_DIR}/samp-plugin-sdk
