# Build stage
FROM node:18-alpine AS builder
WORKDIR /app

# Copy dependency files and install
COPY package*.json ./
RUN npm install

# Copy source code and build
COPY . .
RUN npm run build

# Serve stage
FROM nginx:alpine
# Copy built assets to Nginx html directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port (Nginx defaults to 80)
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
