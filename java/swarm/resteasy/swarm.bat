ECHO OFF
SET port=%1

IF "%1"=="" (
   SET port=8080
)

ECHO ON
call java -Xms1024m -Xmx1024m -Dswarm.http.port=%port% -jar wildfly-swarm-0.0.1-SNAPSHOT-swarm.jar

