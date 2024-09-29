# Respuestas-Prueba-SOFTWARE-DEVELOPER
Respuestas-Prueba-SOFTWARE DEVELOPER


1. Problemas de diseño:

El cuello de botella está en la base de datos, que no admite transacciones simultáneas. Una posible solución sería implementar un sistema de colas de trabajo. Las solicitudes se almacenarían temporalmente en una cola, y el sistema las procesaría de manera secuencial. Esto evitaría sobrecargar la base de datos con múltiples conexiones al mismo tiempo.
Otra alternativa es optimizar las consultas en la base de datos mediante mejoras en los índices o ajustes en el esquema. Además, se podría emplear una base de datos caché como Redis, para reducir la carga causada por consultas repetitivas.


2. Encripta tu mensaje
ver el archivo encripta_tu_mensaje.go contiene el código respuesta en Go
 
4. Suma a cero
ver el archivo suma_a_cero.go contiene el código respuesta en Go
   
5. Aprendizaje

 			Análisis de la aplicación:

			1- Client Apps (Aplicaciones Cliente)

			Las aplicaciones cliente son las interfaces que los usuarios utilizan para interactuar con el sistema. En este caso, hay dos tipos de aplicaciones:

			Web: Aplicación que se accede a través de un navegador.
			Mobile: Aplicación que se utiliza en dispositivos móviles.

			Ambas aplicaciones envían solicitudes al sistema a través del API Gateway.


			2- API Gateway

			El API Gateway actúa como un punto de entrada único para todas las solicitudes provenientes de las aplicaciones cliente.

			Sus funciones clave son:

			Autenticación y autorización: Verifica si el usuario tiene permiso para acceder a ciertos recursos.
			Enrutamiento: Determina qué microservicio debe manejar la solicitud, dirigiéndola al servicio correcto.
			Gestión de cargas: Controla la cantidad de solicitudes para evitar sobrecargas en los microservicios.


			3- Microservicios
			Cada microservicio es un componente independiente que se encarga de una funcionalidad específica del sistema. En este caso, los microservicios mencionados son:

			Catalog DB: Maneja la base de datos que contiene los productos o servicios ofrecidos. Las aplicaciones cliente interactúan con este microservicio para obtener información del catálogo.
			Shopping Cart DB: Se encarga de las operaciones relacionadas con los carritos de compra de los usuarios. Almacena y gestiona los productos agregados por los clientes.
			Discount DB: Aplica y administra los descuentos que se ofrecen a los usuarios, como cupones o promociones.
			Ordering DB: Gestiona el procesamiento de pedidos, incluyendo el flujo desde la creación del pedido hasta su confirmación.

			4- Message Broker (Intermediario de Mensajes)

			El Message Broker actúa como un intermediario que facilita la comunicación entre los microservicios, ayudando a que se comuniquen de forma asíncrona. En este caso:

			Microservicios que envían información al Message Broker:
			Catalog DB y Shopping Cart DB envían eventos al intermediario para notificar cambios, como la actualización de un catálogo o la modificación de un carrito de compras.

			Microservicios que reciben información del Message Broker:
			Discount DB y Ordering DB reciben estos eventos para tomar decisiones basadas en la información actualizada. Por ejemplo, si el catálogo cambia, el sistema de descuentos puede necesitar actualizarse para reflejar nuevas promociones, o el sistema de órdenes puede procesar un pedido basado en la información del carrito.

			Mejoras que consideraría:

			Cache adicional (como Redis) entre el API Gateway y los microservicios:
			Esto podría reducir la carga en los microservicios más consultados, como el Catalog DB, al almacenar en memoria datos que no cambian con frecuencia, mejorando así el rendimiento.

			Escalabilidad del Message Broker:
			Asegurarse de que el Message Broker pueda escalar adecuadamente si la cantidad de mensajes aumenta, ya que se convierte en un punto crítico en la comunicación entre microservicios.

			Logging y monitoreo:
			Implementaría un sistema de logging y monitoreo centralizado para seguir el comportamiento de cada microservicio y el Message Broker, para detectar cuellos de botella o fallas en tiempo real.
   
6. Demostrando tus hallazgos

Presentación de nueva tecnología

Hoy quiero compartirles una nueva tecnología que optimizará el rendimiento de nuestro sistema de órdenes. 
Se trata de los microservicios, una forma de hacer que el sistema sea más escalable y que responda mucho más rápido.
Imaginemos que nuestro sistema actual es como un restaurante de comida rápida donde un solo empleado toma el pedido, cocina, empaqueta y cobra al cliente. Si llegan muchos clientes al mismo tiempo, este empleado no puede hacer todo de manera eficiente, y se forman largas filas. 
Con los microservicios, es como si cada tarea en el restaurante estuviera a cargo de un empleado diferente: uno toma los pedidos, otro cocina, otro se encarga del empaque, y otro cobra. Cada uno hace su parte sin esperar a los demás, lo que hace que todo funcione más rápido. 
Si un área se satura, las demás siguen operando sin problemas. De esta manera, nuestro sistema será mucho más rápido y podrá manejar un mayor volumen de órdenes sin colapsar.
 
