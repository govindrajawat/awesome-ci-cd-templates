package com.example.demo.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.actuate.health.Health;
import org.springframework.boot.actuate.health.HealthIndicator;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.sql.DataSource;
import java.sql.Connection;
import java.sql.SQLException;
import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.Map;

@RestController
public class HealthController implements HealthIndicator {

    @Autowired
    private DataSource dataSource;

    @GetMapping("/health")
    public Map<String, Object> health() {
        Map<String, Object> health = new HashMap<>();
        health.put("status", "UP");
        health.put("timestamp", LocalDateTime.now().toString());
        health.put("database", checkDatabaseHealth());
        return health;
    }

    @Override
    public Health health() {
        try {
            if (checkDatabaseHealth().equals("UP")) {
                return Health.up()
                    .withDetail("database", "UP")
                    .withDetail("timestamp", LocalDateTime.now().toString())
                    .build();
            } else {
                return Health.down()
                    .withDetail("database", "DOWN")
                    .withDetail("timestamp", LocalDateTime.now().toString())
                    .build();
            }
        } catch (Exception e) {
            return Health.down()
                .withDetail("error", e.getMessage())
                .withDetail("timestamp", LocalDateTime.now().toString())
                .build();
        }
    }

    private String checkDatabaseHealth() {
        try (Connection connection = dataSource.getConnection()) {
            if (connection.isValid(5)) {
                return "UP";
            } else {
                return "DOWN";
            }
        } catch (SQLException e) {
            return "DOWN";
        }
    }
} 