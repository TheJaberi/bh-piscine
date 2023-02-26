#!/bin/bash
curl -s https://learn.reboot01.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"ajaberi\"}}){id}}"}'