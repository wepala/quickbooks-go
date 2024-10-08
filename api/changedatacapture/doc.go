// This file was auto-generated by Fern from our API Definition.

// https://developer.intuit.com/docs/api/accounting/changedatacapture
//
// The change data capture (cdc) operation returns a list of objects that have changed since a specified time. This operation is for an app that periodically polls data services in order to refresh its local copy of object data. The app calls the cdc operation, specifying a comma separated list of object types and a date-time stamp specifying how far to look back. Data services returns all objects specified by entityList that have changed since the specified date-time. Look-back time can be up to 30 days.
package changedatacapture
