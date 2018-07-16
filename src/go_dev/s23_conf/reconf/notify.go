package reconf

// 通知应用程序文件改变


type Notifyer  interface {
	Callback(*Config)
}