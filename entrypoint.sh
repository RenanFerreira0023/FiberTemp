#!/bin/sh

# Verifica as permissões do diretório /Logs
echo "Verificando permissões do diretório /Logs..."
ls -ld /Logs
ls -l /Logs

# Inicia o servidor
exec /server
