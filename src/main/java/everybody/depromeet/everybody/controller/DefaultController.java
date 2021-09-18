package everybody.depromeet.everybody.controller;

import everybody.depromeet.everybody.model.User;
import everybody.depromeet.everybody.repository.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequiredArgsConstructor
public class DefaultController {
    private final UserRepository userRepository;

    @GetMapping("/")
    @ResponseStatus(code=HttpStatus.OK)
    public String ping(){
        return "pong-1";
    }

    @GetMapping("/users")
    @ResponseStatus(code=HttpStatus.OK)
    public List<User> listUsers(){
        return userRepository.findAll();
    }


}
