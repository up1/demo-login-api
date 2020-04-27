package com.example.demo.login;

import org.springframework.data.repository.CrudRepository;

import java.util.Optional;

public interface UserReposiotry extends CrudRepository<User, Integer> {

    Optional<User> findByUsernameAndPassword(String username, String password);

}
