// Auto-generated by avdl-compiler v1.2.0 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/notify_ctl.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type NotificationChannels struct {
	Session  bool `codec:"session" json:"session"`
	Users    bool `codec:"users" json:"users"`
	Kbfs     bool `codec:"kbfs" json:"kbfs"`
	Tracking bool `codec:"tracking" json:"tracking"`
}

type SetNotificationsArg struct {
	Channels NotificationChannels `codec:"channels" json:"channels"`
}

type NotifyCtlInterface interface {
	SetNotifications(context.Context, NotificationChannels) error
}

func NotifyCtlProtocol(i NotifyCtlInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.notifyCtl",
		Methods: map[string]rpc.ServeHandlerDescription{
			"setNotifications": {
				MakeArg: func() interface{} {
					ret := make([]SetNotificationsArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SetNotificationsArg)
					if !ok {
						err = rpc.NewTypeError((*[]SetNotificationsArg)(nil), args)
						return
					}
					err = i.SetNotifications(ctx, (*typedArgs)[0].Channels)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type NotifyCtlClient struct {
	Cli rpc.GenericClient
}

func (c NotifyCtlClient) SetNotifications(ctx context.Context, channels NotificationChannels) (err error) {
	__arg := SetNotificationsArg{Channels: channels}
	err = c.Cli.Call(ctx, "keybase.1.notifyCtl.setNotifications", []interface{}{__arg}, nil)
	return
}