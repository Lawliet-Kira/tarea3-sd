### Tarea 3 - Sistema Distribuido
### GRUPO 32
#### Nombres + Rol:
*  Diego Reinoso   + 201523518-K
*  David Monsalves + 201773547-3
*  Nicolas Silva   + 201710516-7
    
Para la ejecución de la tarea, es necesario desplegar 6 terminales; una para cada entidad a ejecutar:

*   Informante A. Tano      - Maquina dist 125
*   Informante A. Thrawn    - Maquina dist 126
*   Broker                  - Maquina dist 128
*   Leia                    - Maquina dist 127
*   Servidor Fulcrum 1      - Maquina dist 125
*   Servidor Fulcrum 2      - Maquina dist 126 (Servidor dominante)
*   Servidor Fulcrum 3      - Maquina dist 127

Ubicar la ruta actual de cada terminal en el directorio "tarea3-sd". Posteriormente, ejecutar las siguientes 
sentencias sobre cada terminal ,respectivamente, y en el orden correspondiente:

*   Terminal Broker                 -->    $ make broker
*   Terminal Servidor Fulcrum 1     -->    $ make fulcrum
*   Terminal Servidor Fulcrum 3     -->    $ make fulcrum
*   Terminal Servidor Fulcrum 2     -->    $ make fulcrum
*   Terminal Informante A. Tano     -->    $ make info
*   Terminal Informante A. Thrawn   -->    $ make info 
*   Terminal Leia                   -->    $ make leia

La interacción con la tarea se debe realizar en las terminales de Informante y Leia, las cuales desplegaran un menú donde
se deberá escojer una de las acciones a realizar. A modo de ejemplo, el menú de un Informante corresponde a:

Ingrese Operación: 
	1. Añadir Ciudad
	2. Actualizar Nombre
	3. Actualiar Numero
	4. Eliminar Ciudad

Luego de escojer una opción se debe ingresar el comando correspondiente a la acción escogida. A modo de ejemplo, para el caso
de añadir una ciduad: 

Ingresar Comando:
Tierra Santiago 122

La sintaxis para el ingreso de los comandos según la acción escogida de resumen a continuación:
    
            ACCION                             COMANDO

*       Añadir Ciudad       |     nombre_planeta nombre_ciudad [nuevo_valor]
*       Actualizar Nombre   |     nombre_planeta nombre_ciudad nuevo_valor
*       Actualizar Numero   |     nombre_planeta nombre_ciudad nuevo valor
*       Eliminar Ciudad     |     nombre_planeta nombre_ciudad












