package com.example.demo.login;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class LoginExceptionHandler {

    @ExceptionHandler({UserNotFoundException.class})
    @ResponseStatus(value = HttpStatus.OK)
    @ResponseBody
    public String handleCustomException(UserNotFoundException ex) {
        return ex.getMessage();
    }
}
