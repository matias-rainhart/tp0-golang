package main

import (
    "bufio"
    "client/globals"
    "client/utils"
    "log"
    "os"
)

func main() {
    utils.ConfigurarLogger()

    //loggear "Hola soy un log" usando la biblioteca log
    log.Println("Hola soy un log")

    // validar que la config este cargada correctamente
    globals.ClientConfig = utils.IniciarConfiguracion("config.json")
    if globals.ClientConfig == nil {
        log.Fatalf("No se pudo cargar la configuración")
    }

    // loggeamos el valor de la config
    log.Println(globals.ClientConfig.Mensaje)

	
	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él
	
	// enviar un mensaje al servidor con el valor de la config

	// leer de la consola el mensaje
	// existe utils.LeerConsola()
	var text string
    var mensajesConsola []string
    for text != "\n" {
        reader := bufio.NewReader(os.Stdin)
        text, _ = reader.ReadString('\n')
        log.Print(text)
        mensajesConsola = append(mensajesConsola, text)
    }

    // enviar el mensaje al servidor
    ip := globals.ClientConfig.Ip
	puerto := globals.ClientConfig.Puerto
    clave := globals.ClientConfig.Clave
    utils.EnviarMensaje(ip, puerto, clave)

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete(mensajesConsola, ip, puerto)
}
