package iwd

import (
	"github.com/godbus/dbus/v5"
)

const (
	objectStation                 = "net.connman.iwd.Station"
	callStationScan               = "net.connman.iwd.Station.Scan"
	callStationDisconnect         = "net.connman.iwd.Station.Disconnect"
	callStationGetOrderedNetworks = "net.connman.iwd.Station.GetOrderedNetworks"
)

// Station refers to net.connman.iwd.Station
type Station struct {
	Path             dbus.ObjectPath
	ConnectedNetwork dbus.ObjectPath
	Scanning         bool
	State            string
}

// Scan scans for wireless networks
func (s *Station) Scan(conn *dbus.Conn) error {
	obj := conn.Object(objectIwd, s.Path)
	call := obj.Call(callStationScan, 0)
	if call.Err != nil {
		return call.Err
	}

	return nil
}

// Disconnect station
func (s *Station) Disconnect(conn *dbus.Conn) error {
	obj := conn.Object(objectIwd, s.Path)
	call := obj.Call(callStationDisconnect, 0)
	if call.Err != nil {
		return call.Err
	}

	return nil
}
