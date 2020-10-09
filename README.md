# Tarea1Dist

Conexion Cliente (Máquina 10.10.28.155) y Logística (máquina 10.10.28.154) casi lista
  - El servidor de logística está en grpc/server, tiene que abrirse con 'go run server.go'
  - El cliente es el que está en la carpeta principal, se tiene que abrir con 'go run Cliente.go'
  - Solo falta la opcion 3 del cliente, que se supone que tiene que enviar el estado del pedido de acuerdo al id dado.
  
 En la carpeta Pruebas está el archivo 'vercsv.go'. Lo hice para poder ver la info que está en cada csv.
 
 Todos los archivos csv están en la carpeta archivos. El programa lee desde retail.csv y pymes.csv. El archivo indexAct lo usa para ver el ID que le da a cada pedido. Al correr el Cliente se crea 'results.csv' en el servidor, que indica las órdenes que han sido ingresadas por el cliente.
