package bishop

import (
	"unsafe"
)

// #cgo LDFLAGS: -ldl
import "C"

//export Supports
func Supports() C.uint {
	return C.uint(ISupports())
}

//export Load
func Load(ppData **C.void) bool {
	return ILoad((*uintptr)(unsafe.Pointer(ppData)), 0)
}

//export Unload
func Unload() {
	IUnload(0)
}

//export ProcessTick
func ProcessTick() {
	IProcessTick(0)
}

var CallbackOnGameModeInit = func( ) bool { return true }

//export OnGameModeInit
func OnGameModeInit( ) bool {
	return CallbackOnGameModeInit()
}

var CallbackOnGameModeExit = func( ) bool { return true }

//export OnGameModeExit
func OnGameModeExit( ) bool {
	return CallbackOnGameModeExit()
}

var CallbackOnPlayerConnect = func(playerid int) bool { return true }

//export OnPlayerConnect
func OnPlayerConnect(playerid C.int) bool {
	return CallbackOnPlayerConnect(int(playerid))
}

var CallbackOnPlayerDisconnect = func(playerid int, reason int) bool { return true }

//export OnPlayerDisconnect
func OnPlayerDisconnect(playerid C.int, reason C.int) bool {
	return CallbackOnPlayerDisconnect(int(playerid), int(reason))
}

var CallbackOnPlayerSpawn = func(playerid int) bool { return true }

//export OnPlayerSpawn
func OnPlayerSpawn(playerid C.int) bool {
	return CallbackOnPlayerSpawn(int(playerid))
}

var CallbackOnPlayerDeath = func(playerid int, killerid int, reason int) bool { return true }

//export OnPlayerDeath
func OnPlayerDeath(playerid C.int, killerid C.int, reason C.int) bool {
	return CallbackOnPlayerDeath(int(playerid), int(killerid), int(reason))
}

var CallbackOnVehicleSpawn = func(vehicleid int) bool { return true }

//export OnVehicleSpawn
func OnVehicleSpawn(vehicleid C.int) bool {
	return CallbackOnVehicleSpawn(int(vehicleid))
}

var CallbackOnVehicleDeath = func(vehicleid int, killerid int) bool { return true }

//export OnVehicleDeath
func OnVehicleDeath(vehicleid C.int, killerid C.int) bool {
	return CallbackOnVehicleDeath(int(vehicleid), int(killerid))
}

var CallbackOnPlayerText = func(playerid int, text string) bool { return true }

//export OnPlayerText
func OnPlayerText(playerid C.int, text *C.char) bool {
	return CallbackOnPlayerText(int(playerid), C.GoString(text))
}

var CallbackOnPlayerCommandText = func(playerid int, cmdtext string) bool { return true }

//export OnPlayerCommandText
func OnPlayerCommandText(playerid C.int, cmdtext *C.char) bool {
	return CallbackOnPlayerCommandText(int(playerid), C.GoString(cmdtext))
}

var CallbackOnPlayerRequestClass = func(playerid int, classid int) bool { return true }

//export OnPlayerRequestClass
func OnPlayerRequestClass(playerid C.int, classid C.int) bool {
	return CallbackOnPlayerRequestClass(int(playerid), int(classid))
}

var CallbackOnPlayerEnterVehicle = func(playerid int, vehicleid int, ispassenger bool) bool { return true }

//export OnPlayerEnterVehicle
func OnPlayerEnterVehicle(playerid C.int, vehicleid C.int, ispassenger C.int) bool {
	return CallbackOnPlayerEnterVehicle(int(playerid), int(vehicleid), int(ispassenger) != 0)
}

var CallbackOnPlayerExitVehicle = func(playerid int, vehicleid int) bool { return true }

//export OnPlayerExitVehicle
func OnPlayerExitVehicle(playerid C.int, vehicleid C.int) bool {
	return CallbackOnPlayerExitVehicle(int(playerid), int(vehicleid))
}

var CallbackOnPlayerStateChange = func(playerid int, newstate int, oldstate int) bool { return true }

//export OnPlayerStateChange
func OnPlayerStateChange(playerid C.int, newstate C.int, oldstate C.int) bool {
	return CallbackOnPlayerStateChange(int(playerid), int(newstate), int(oldstate))
}

var CallbackOnPlayerEnterCheckpoint = func(playerid int) bool { return true }

//export OnPlayerEnterCheckpoint
func OnPlayerEnterCheckpoint(playerid C.int) bool {
	return CallbackOnPlayerEnterCheckpoint(int(playerid))
}

var CallbackOnPlayerLeaveCheckpoint = func(playerid int) bool { return true }

//export OnPlayerLeaveCheckpoint
func OnPlayerLeaveCheckpoint(playerid C.int) bool {
	return CallbackOnPlayerLeaveCheckpoint(int(playerid))
}

var CallbackOnPlayerEnterRaceCheckpoint = func(playerid int) bool { return true }

//export OnPlayerEnterRaceCheckpoint
func OnPlayerEnterRaceCheckpoint(playerid C.int) bool {
	return CallbackOnPlayerEnterRaceCheckpoint(int(playerid))
}

var CallbackOnPlayerLeaveRaceCheckpoint = func(playerid int) bool { return true }

//export OnPlayerLeaveRaceCheckpoint
func OnPlayerLeaveRaceCheckpoint(playerid C.int) bool {
	return CallbackOnPlayerLeaveRaceCheckpoint(int(playerid))
}

var CallbackOnRconCommand = func(cmd string) bool { return true }

//export OnRconCommand
func OnRconCommand(cmd *C.char) bool {
	return CallbackOnRconCommand(C.GoString(cmd))
}

var CallbackOnPlayerRequestSpawn = func(playerid int) bool { return true }

//export OnPlayerRequestSpawn
func OnPlayerRequestSpawn(playerid C.int) bool {
	return CallbackOnPlayerRequestSpawn(int(playerid))
}

var CallbackOnObjectMoved = func(objectid int) bool { return true }

//export OnObjectMoved
func OnObjectMoved(objectid C.int) bool {
	return CallbackOnObjectMoved(int(objectid))
}

var CallbackOnPlayerObjectMoved = func(playerid int, objectid int) bool { return true }

//export OnPlayerObjectMoved
func OnPlayerObjectMoved(playerid C.int, objectid C.int) bool {
	return CallbackOnPlayerObjectMoved(int(playerid), int(objectid))
}

var CallbackOnPlayerPickUpPickup = func(playerid int, pickupid int) bool { return true }

//export OnPlayerPickUpPickup
func OnPlayerPickUpPickup(playerid C.int, pickupid C.int) bool {
	return CallbackOnPlayerPickUpPickup(int(playerid), int(pickupid))
}

var CallbackOnVehicleMod = func(playerid int, vehicleid int, componentid int) bool { return true }

//export OnVehicleMod
func OnVehicleMod(playerid C.int, vehicleid C.int, componentid C.int) bool {
	return CallbackOnVehicleMod(int(playerid), int(vehicleid), int(componentid))
}

var CallbackOnEnterExitModShop = func(playerid int, enterexit bool, interiorid int) bool { return true }

//export OnEnterExitModShop
func OnEnterExitModShop(playerid C.int, enterexit C.int, interiorid C.int) bool {
	return CallbackOnEnterExitModShop(int(playerid), int(enterexit) != 0, int(interiorid))
}

var CallbackOnVehiclePaintjob = func(playerid int, vehicleid int, paintjobid int) bool { return true }

//export OnVehiclePaintjob
func OnVehiclePaintjob(playerid C.int, vehicleid C.int, paintjobid C.int) bool {
	return CallbackOnVehiclePaintjob(int(playerid), int(vehicleid), int(paintjobid))
}

var CallbackOnVehicleRespray = func(playerid int, vehicleid int, color1 int, color2 int) bool { return true }

//export OnVehicleRespray
func OnVehicleRespray(playerid C.int, vehicleid C.int, color1 C.int, color2 C.int) bool {
	return CallbackOnVehicleRespray(int(playerid), int(vehicleid), int(color1), int(color2))
}

var CallbackOnVehicleDamageStatusUpdate = func(vehicleid int, playerid int) bool { return true }

//export OnVehicleDamageStatusUpdate
func OnVehicleDamageStatusUpdate(vehicleid C.int, playerid C.int) bool {
	return CallbackOnVehicleDamageStatusUpdate(int(vehicleid), int(playerid))
}

var CallbackOnUnoccupiedVehicleUpdate = func(vehicleid int, playerid int, passenger_seat int, new_x float32, new_y float32, new_z float32, vel_x float32, vel_y float32, vel_z float32) bool { return true }

//export OnUnoccupiedVehicleUpdate
func OnUnoccupiedVehicleUpdate(vehicleid C.int, playerid C.int, passenger_seat C.int, new_x C.float, new_y C.float, new_z C.float, vel_x C.float, vel_y C.float, vel_z C.float) bool {
	return CallbackOnUnoccupiedVehicleUpdate(int(vehicleid), int(playerid), int(passenger_seat), float32(new_x), float32(new_y), float32(new_z), float32(vel_x), float32(vel_y), float32(vel_z))
}

var CallbackOnPlayerSelectedMenuRow = func(playerid int, row int) bool { return true }

//export OnPlayerSelectedMenuRow
func OnPlayerSelectedMenuRow(playerid C.int, row C.int) bool {
	return CallbackOnPlayerSelectedMenuRow(int(playerid), int(row))
}

var CallbackOnPlayerExitedMenu = func(playerid int) bool { return true }

//export OnPlayerExitedMenu
func OnPlayerExitedMenu(playerid C.int) bool {
	return CallbackOnPlayerExitedMenu(int(playerid))
}

var CallbackOnPlayerInteriorChange = func(playerid int, newinteriorid int, oldinteriorid int) bool { return true }

//export OnPlayerInteriorChange
func OnPlayerInteriorChange(playerid C.int, newinteriorid C.int, oldinteriorid C.int) bool {
	return CallbackOnPlayerInteriorChange(int(playerid), int(newinteriorid), int(oldinteriorid))
}

var CallbackOnPlayerKeyStateChange = func(playerid int, newkeys int, oldkeys int) bool { return true }

//export OnPlayerKeyStateChange
func OnPlayerKeyStateChange(playerid C.int, newkeys C.int, oldkeys C.int) bool {
	return CallbackOnPlayerKeyStateChange(int(playerid), int(newkeys), int(oldkeys))
}

var CallbackOnRconLoginAttempt = func(ip string, password string, success bool) bool { return true }

//export OnRconLoginAttempt
func OnRconLoginAttempt(ip *C.char, password *C.char, success C.int) bool {
	return CallbackOnRconLoginAttempt(C.GoString(ip), C.GoString(password), int(success) != 0)
}

var CallbackOnPlayerUpdate = func(playerid int) bool { return true }

//export OnPlayerUpdate
func OnPlayerUpdate(playerid C.int) bool {
	return CallbackOnPlayerUpdate(int(playerid))
}

var CallbackOnPlayerStreamIn = func(playerid int, forplayerid int) bool { return true }

//export OnPlayerStreamIn
func OnPlayerStreamIn(playerid C.int, forplayerid C.int) bool {
	return CallbackOnPlayerStreamIn(int(playerid), int(forplayerid))
}

var CallbackOnPlayerStreamOut = func(playerid int, forplayerid int) bool { return true }

//export OnPlayerStreamOut
func OnPlayerStreamOut(playerid C.int, forplayerid C.int) bool {
	return CallbackOnPlayerStreamOut(int(playerid), int(forplayerid))
}

var CallbackOnVehicleStreamIn = func(vehicleid int, forplayerid int) bool { return true }

//export OnVehicleStreamIn
func OnVehicleStreamIn(vehicleid C.int, forplayerid C.int) bool {
	return CallbackOnVehicleStreamIn(int(vehicleid), int(forplayerid))
}

var CallbackOnVehicleStreamOut = func(vehicleid int, forplayerid int) bool { return true }

//export OnVehicleStreamOut
func OnVehicleStreamOut(vehicleid C.int, forplayerid C.int) bool {
	return CallbackOnVehicleStreamOut(int(vehicleid), int(forplayerid))
}

var CallbackOnActorStreamIn = func(actorid int, forplayerid int) bool { return true }

//export OnActorStreamIn
func OnActorStreamIn(actorid C.int, forplayerid C.int) bool {
	return CallbackOnActorStreamIn(int(actorid), int(forplayerid))
}

var CallbackOnActorStreamOut = func(actorid int, forplayerid int) bool { return true }

//export OnActorStreamOut
func OnActorStreamOut(actorid C.int, forplayerid C.int) bool {
	return CallbackOnActorStreamOut(int(actorid), int(forplayerid))
}

var CallbackOnDialogResponse = func(playerid int, dialogid int, response int, listitem int, inputtext string) bool { return true }

//export OnDialogResponse
func OnDialogResponse(playerid C.int, dialogid C.int, response C.int, listitem C.int, inputtext *C.char) bool {
	return CallbackOnDialogResponse(int(playerid), int(dialogid), int(response), int(listitem), C.GoString(inputtext))
}

var CallbackOnPlayerTakeDamage = func(playerid int, issuerid int, amount float32, weaponid int, bodypart int) bool { return true }

//export OnPlayerTakeDamage
func OnPlayerTakeDamage(playerid C.int, issuerid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	return CallbackOnPlayerTakeDamage(int(playerid), int(issuerid), float32(amount), int(weaponid), int(bodypart))
}

var CallbackOnPlayerGiveDamage = func(playerid int, damagedid int, amount float32, weaponid int, bodypart int) bool { return true }

//export OnPlayerGiveDamage
func OnPlayerGiveDamage(playerid C.int, damagedid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	return CallbackOnPlayerGiveDamage(int(playerid), int(damagedid), float32(amount), int(weaponid), int(bodypart))
}

var CallbackOnPlayerGiveDamageActor = func(playerid int, damaged_actorid int, amount float32, weaponid int, bodypart int) bool { return true }

//export OnPlayerGiveDamageActor
func OnPlayerGiveDamageActor(playerid C.int, damaged_actorid C.int, amount C.float, weaponid C.int, bodypart C.int) bool {
	return CallbackOnPlayerGiveDamageActor(int(playerid), int(damaged_actorid), float32(amount), int(weaponid), int(bodypart))
}

var CallbackOnPlayerClickMap = func(playerid int, fX float32, fY float32, fZ float32) bool { return true }

//export OnPlayerClickMap
func OnPlayerClickMap(playerid C.int, fX C.float, fY C.float, fZ C.float) bool {
	return CallbackOnPlayerClickMap(int(playerid), float32(fX), float32(fY), float32(fZ))
}

var CallbackOnPlayerClickTextDraw = func(playerid int, clickedid int) bool { return true }

//export OnPlayerClickTextDraw
func OnPlayerClickTextDraw(playerid C.int, clickedid C.int) bool {
	return CallbackOnPlayerClickTextDraw(int(playerid), int(clickedid))
}

var CallbackOnPlayerClickPlayerTextDraw = func(playerid int, playertextid int) bool { return true }

//export OnPlayerClickPlayerTextDraw
func OnPlayerClickPlayerTextDraw(playerid C.int, playertextid C.int) bool {
	return CallbackOnPlayerClickPlayerTextDraw(int(playerid), int(playertextid))
}

var CallbackOnIncomingConnection = func(playerid int, ip_address string, port int) bool { return true }

//export OnIncomingConnection
func OnIncomingConnection(playerid C.int, ip_address *C.char, port C.int) bool {
	return CallbackOnIncomingConnection(int(playerid), C.GoString(ip_address), int(port))
}

var CallbackOnTrailerUpdate = func(playerid int, vehicleid int) bool { return true }

//export OnTrailerUpdate
func OnTrailerUpdate(playerid C.int, vehicleid C.int) bool {
	return CallbackOnTrailerUpdate(int(playerid), int(vehicleid))
}

var CallbackOnVehicleSirenStateChange = func(playerid int, vehicleid int, newstate int) bool { return true }

//export OnVehicleSirenStateChange
func OnVehicleSirenStateChange(playerid C.int, vehicleid C.int, newstate C.int) bool {
	return CallbackOnVehicleSirenStateChange(int(playerid), int(vehicleid), int(newstate))
}

var CallbackOnPlayerClickPlayer = func(playerid int, clickedplayerid int, source int) bool { return true }

//export OnPlayerClickPlayer
func OnPlayerClickPlayer(playerid C.int, clickedplayerid C.int, source C.int) bool {
	return CallbackOnPlayerClickPlayer(int(playerid), int(clickedplayerid), int(source))
}

var CallbackOnPlayerEditObject = func(playerid int, playerobject bool, objectid int, response int, fX float32, fY float32, fZ float32, fRotX float32, fRotY float32, fRotZ float32) bool { return true }

//export OnPlayerEditObject
func OnPlayerEditObject(playerid C.int, playerobject C.int, objectid C.int, response C.int, fX C.float, fY C.float, fZ C.float, fRotX C.float, fRotY C.float, fRotZ C.float) bool {
	return CallbackOnPlayerEditObject(int(playerid), int(playerobject) != 0, int(objectid), int(response), float32(fX), float32(fY), float32(fZ), float32(fRotX), float32(fRotY), float32(fRotZ))
}

var CallbackOnPlayerEditAttachedObject = func(playerid int, response int, index int, modelid int, boneid int, fOffsetX float32, fOffsetY float32, fOffsetZ float32, fRotX float32, fRotY float32, fRotZ float32, fScaleX float32, fScaleY float32, fScaleZ float32) bool { return true }

//export OnPlayerEditAttachedObject
func OnPlayerEditAttachedObject(playerid C.int, response C.int, index C.int, modelid C.int, boneid C.int, fOffsetX C.float, fOffsetY C.float, fOffsetZ C.float, fRotX C.float, fRotY C.float, fRotZ C.float, fScaleX C.float, fScaleY C.float, fScaleZ C.float) bool {
	return CallbackOnPlayerEditAttachedObject(int(playerid), int(response), int(index), int(modelid), int(boneid), float32(fOffsetX), float32(fOffsetY), float32(fOffsetZ), float32(fRotX), float32(fRotY), float32(fRotZ), float32(fScaleX), float32(fScaleY), float32(fScaleZ))
}

var CallbackOnPlayerSelectObject = func(playerid int, t int, objectid int, modelid int, fX float32, fY float32, fZ float32) bool { return true }

//export OnPlayerSelectObject
func OnPlayerSelectObject(playerid C.int, t C.int, objectid C.int, modelid C.int, fX C.float, fY C.float, fZ C.float) bool {
	return CallbackOnPlayerSelectObject(int(playerid), int(t), int(objectid), int(modelid), float32(fX), float32(fY), float32(fZ))
}

var CallbackOnPlayerWeaponShot = func(playerid int, weaponid int, hittype int, hitid int, fX float32, fY float32, fZ float32) bool { return true }

//export OnPlayerWeaponShot
func OnPlayerWeaponShot(playerid C.int, weaponid C.int, hittype C.int, hitid C.int, fX C.float, fY C.float, fZ C.float) bool {
	return CallbackOnPlayerWeaponShot(int(playerid), int(weaponid), int(hittype), int(hitid), float32(fX), float32(fY), float32(fZ))
}

var CallbackOnPlayerRequestDownload = func(playerid int, t int, crc int) bool { return true }

//export OnPlayerRequestDownload
func OnPlayerRequestDownload(playerid C.int, t C.int, crc C.int) bool {
	return CallbackOnPlayerRequestDownload(int(playerid), int(t), int(crc))
}

var CallbackOnHTTPResponse = func(index int, response_code int, data string) { }

//export OnHTTPResponse
func OnHTTPResponse(index C.int, response_code C.int, data *C.char) {
	CallbackOnHTTPResponse(int(index), int(response_code), C.GoString(data))
}

