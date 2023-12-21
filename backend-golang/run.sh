#!/bin/bash

# Este script ejecuta el archivo main desde el path absoluto

# Establecer el directorio de trabajo
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$DIR"

ls -la /backend-golang

# Ejecutar el archivo main
./finalApp