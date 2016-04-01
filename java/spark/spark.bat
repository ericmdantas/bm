ECHO OFF
SET port=%1

IF "%1"=="" (
   SET port=8080
)

ECHO ON
call java -Dport=%port% -jar target/spark-rest-0.0.1-SNAPSHOT.jar

pause
