import { RoomTypesInterface } from "./IRoomTypes";
import { RoomZonesInterface } from "./IRoomZones";
import {AdminInterface} from "./IAdmins";

export interface RoomsInterface {
  ID?: number,
  RoomNumber: string,
  RoomZoneID?: number,
  RoomZone?: RoomZonesInterface,
  RoomTypeID?: number,
  RoomType?: RoomTypesInterface,
  Admin?: string,
  AdminID?: AdminInterface
}
