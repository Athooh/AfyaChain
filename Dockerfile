# Use the official MySQL image from the Docker Hub
FROM mysql:8.0

# Set environment variables for the MySQL database
ENV MYSQL_ROOT_PASSWORD=root_password
ENV MYSQL_DATABASE=afya_chain_db
ENV MYSQL_USER=new_username
ENV MYSQL_PASSWORD=new_password

# Copy the initialization SQL script to the Docker image
COPY init.sql /docker-entrypoint-initdb.d/

# Expose the MySQL port
EXPOSE 3306
