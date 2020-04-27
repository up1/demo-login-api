package com.example.demo.login;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.Optional;

@RestController
public class LoginController {

    private UserReposiotry userReposiotry;

    @Autowired
    public LoginController(UserReposiotry userReposiotry) {
        this.userReposiotry = userReposiotry;
    }

    @PostMapping("/api/v1/login")
    public LoginResponse login(@RequestBody LoginRequest request) {
        Optional<User> user =  userReposiotry.findByUsernameAndPassword(request.getUsername(), request.getPassword());
        if(!user.isPresent()) {
            throw new UserNotFoundException("User not found with " + request.getUsername());
        }
        User loggedUser = user.get();
        LoginResponse response = new LoginResponse();
        response.setId(loggedUser.getId());
        response.setFirstname(loggedUser.getFirstname());
        response.setLastname(loggedUser.getLastname());
        response.setEmail(loggedUser.getEmail());
        return response;
    }

}
