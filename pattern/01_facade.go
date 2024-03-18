package main

import "fmt"

type powerBattery struct {}

func (p *powerBattery) checked(){
	fmt.Println("battery checked")
}

type android struct {}

func (a *android) start(){
	fmt.Println("android started")
}

func (a *android) stop(){
	fmt.Println("android shuted down")
}

type wifi struct {}

func (w *wifi) onWiFi(){
	fmt.Println("wifi is on")
}

func (w *wifi) offWiFi(){
	fmt.Println("wifi is off")
}

type Facade struct{
	battery *powerBattery
	system *android
	connection *wifi
}

func(f *Facade) NewFacade() *Facade{
	f.battery = &powerBattery{}
	f.system = &android{}
	f.connection = &wifi{}
	return f
}
	
func (f *Facade) Start(){
	f.battery.checked()
	f.system.start()
	f.connection.onWiFi()
}

func (f *Facade) Stop(){
	f.system.stop()
	f.connection.offWiFi()
}

func main(){
	f := &Facade{}
	f.NewFacade()
	f.Start()
	f.Stop()
}