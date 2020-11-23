/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Project API Booking
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntBooking,
    EntBookingFromJSON,
    EntBookingFromJSONTyped,
    EntBookingToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntOperationroomEdges
 */
export interface EntOperationroomEdges {
    /**
     * OperationroomID holds the value of the operationroom_id edge.
     * @type {Array<EntBooking>}
     * @memberof EntOperationroomEdges
     */
    operationroomID?: Array<EntBooking>;
}

export function EntOperationroomEdgesFromJSON(json: any): EntOperationroomEdges {
    return EntOperationroomEdgesFromJSONTyped(json, false);
}

export function EntOperationroomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntOperationroomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'operationroomID': !exists(json, 'operationroomID') ? undefined : ((json['operationroomID'] as Array<any>).map(EntBookingFromJSON)),
    };
}

export function EntOperationroomEdgesToJSON(value?: EntOperationroomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'operationroomID': value.operationroomID === undefined ? undefined : ((value.operationroomID as Array<any>).map(EntBookingToJSON)),
    };
}


