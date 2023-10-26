# NOTAS - GO INTERMEDIO

## TESTING
### COVERAGE
Comandos a tener en cuenta a la hora de construir pruebas unitarias en GO

Ejecutar test unitarios:
```
go test
```
Construir informe de cobertura de pruebas:
```
go test -coverprofile=coverage.out
```
Obtener porcentaje de cobertura de pruebas en consola:
```
go tool cover -func=coverage.out
```
Obtener informe detallado de conertura en HTML:
```
go tool cover -html=coverage.out
```
## PROFILING
Ejecutar profiling de las pruebas unitarias:
```
go test -cpuprofile=cpu.out
```
Revisar informe de profiling en consola:
```
go tool pprof cpu.out
```