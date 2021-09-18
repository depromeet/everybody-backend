package everybody.depromeet.everybody.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.Entity;
import javax.persistence.Id;
import java.io.Serializable;

@Entity
@Getter
@Setter
public class User{
    @Id
    private String id;
    private String description;
}
