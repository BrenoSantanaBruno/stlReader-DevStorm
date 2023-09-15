# Use uma imagem Node.js como base
FROM node:latest

# Defina o diretório de trabalho no contêiner
WORKDIR /app

# Copie o arquivo package.json e package-lock.json (se existir)
COPY package*.json ./

# Instale as dependências
RUN npm install

# Copie todo o código-fonte para o diretório de trabalho no contêiner
COPY . .

# Construa a aplicação React
RUN npm run build

# Exponha a porta em que a aplicação React será executada (geralmente 3000)
EXPOSE 3000

# Inicie a aplicação React quando o contêiner for iniciado
CMD [ "npm", "start" ]