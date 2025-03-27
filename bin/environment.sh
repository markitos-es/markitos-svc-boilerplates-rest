#!/bin/bash

#:[.'.]:> ============================================
#:[.'.]:> 🌍 CONFIGURACIÓN CENTRALIZADA DE ENTORNO
#:[.'.]:> ============================================
#:[.'.]:> Este script centraliza todas las configuraciones
#:[.'.]:> predeterminadas del proyecto para evitar duplicación
#:[.'.]:> y mantener consistencia entre los diferentes scripts.
#:[.'.]:> ============================================

#:[.'.]:> Valores predeterminados para conexión a la base de datos
DEFAULT_DATABASE_HOST="localhost"
DEFAULT_DATABASE_USER="admin"
DEFAULT_DATABASE_PASSWORD="admin"
DEFAULT_DATABASE_NAME="markitos-svc-boilerplates-rest"
DEFAULT_DATABASE_SSL_MODE="disable"

#:[.'.]:> Construir DSN predeterminado
DEFAULT_DATABASE_DSN="host=${DEFAULT_DATABASE_HOST} user=${DEFAULT_DATABASE_USER} password=${DEFAULT_DATABASE_PASSWORD} dbname=${DEFAULT_DATABASE_NAME} sslmode=${DEFAULT_DATABASE_SSL_MODE}"

#:[.'.]:> Otros valores predeterminados
DEFAULT_HTTP_SERVER_ADDRESS=":3000"

#:[.'.]:> Función para configurar variables de entorno
#:[.'.]:> Esta función establece las variables si no están definidas
#:[.'.]:> y las exporta para que estén disponibles para los procesos hijos
function setup_environment() {
    #:[.'.]:> Establecer variables si no están definidas
    : ${DATABASE_DSN:="${DEFAULT_DATABASE_DSN}"}
    : ${HTTP_SERVER_ADDRESS:="${DEFAULT_HTTP_SERVER_ADDRESS}"}

    #:[.'.]:> Exportar variables
    export DATABASE_DSN
    export HTTP_SERVER_ADDRESS
}

#:[.'.]:> Función para mostrar la configuración actual
#:[.'.]:> Parámetro $1 == "full" mostrará todas las variables
#:[.'.]:> Sin parámetros mostrará solo DATABASE_DSN
function show_config() {
    echo "#:[.'.]:> 🚀 Iniciando con configuración:"
    echo "#:[.'.]:> 📊 DATABASE_DSN=$DATABASE_DSN"
    
    if [[ "${1:-}" == "full" ]]; then
        echo "#:[.'.]:> 🌐 HTTP_SERVER_ADDRESS=$HTTP_SERVER_ADDRESS"
    fi
    
    echo "#:[.'.]:>-------------------------------------"
}