package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log")
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente
	
	// loggeamos el valor de la config
	if globals.ClientConfig  == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}
	log.Printf("Configuracion: %+v", globals.ClientConfig)
	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip,globals.ClientConfig.Puerto, "Hola servidor")
	// leer de la consola el mensaje
	for {
        // Leemos de la consola el mensaje
        //mensaje := utils.LeerConsola()

        // Generamos un paquete y lo enviamos al servidor
        utils.GenerarYEnviarPaquete()
    }
}
