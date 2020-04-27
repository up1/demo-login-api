package com.example.demo.login;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.context.SpringBootTest.WebEnvironment;
import org.springframework.boot.test.web.client.TestRestTemplate;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest(webEnvironment = WebEnvironment.RANDOM_PORT)
class LoginControllerTest {

    @Autowired
    TestRestTemplate restTemplate;

    @Autowired
    UserReposiotry userReposiotry;

    @Test
    void login_success_with_username_and_password() {
        // Arrange
        LoginRequest request = new LoginRequest("Somkiat", "PasswordSomKiat");
        userReposiotry.save(new User(1, "Somkiat", "PasswordSomKiat", "somkiat", "Puisung", "somkiat@xxx.com"));
        // Act
        LoginResponse response = restTemplate.postForObject("/api/v1/login", request, LoginResponse.class);
        // Assert
        assertEquals("somkiat", response.getFirstname());
        assertEquals("Puisung", response.getLastname());
        assertEquals("somkiat@xxx.com", response.getEmail());
    }
}