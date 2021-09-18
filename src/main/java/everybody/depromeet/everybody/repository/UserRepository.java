package everybody.depromeet.everybody.repository;

import everybody.depromeet.everybody.model.User;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User, String> {
}
