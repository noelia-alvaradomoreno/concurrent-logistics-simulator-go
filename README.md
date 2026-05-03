## Go Logistics Pipeline - Arquitectura del Sistema
El diseño se basa en un modelo de Pipeline, donde el ciclo de vida del paquete se divide en etapas independientes. Esta estructura permite que cada componente trabaje de forma asíncrona, maximizando el rendimiento del procesador.

*   **Ingesta y Validación:** El paquete sorter actúa como primer filtro, validando destinos y preparando la información clave (ID y estados) del paquete.
*   **Orquestación:** El main gestiona la sincronización mediante canales, asegurando un flujo de datos sin bloqueos innecesarios.
*   **Despacho:** El paquete dispatcher consume los datos finales para ejecutar la salida del sistema.

## Decisiones de Ingeniería

* Se priorizó el uso de Channels sobre el uso de Mutex tradicionales para gestionar el acceso a memoria. Esto no solo simplifica el código, sino que elimina el riesgo de *deadlocks* y condiciones de carrera de manera nativa.
* El main.go funciona estrictamente como un orquestador de infraestructura. Toda la lógica de negocio reside en sus propios paquetes, facilitando la escalabilidad y las pruebas unitarias.
* Implementación de cierres controlados de canales para asegurar que ninguna goroutine quede en espera (evitando fugas de memoria).

## Ejecución

```bash
# Ejecutar el simulador
go run main.go

