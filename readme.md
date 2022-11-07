```
CREATE TABLE public.custom_user (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(30) NOT NULL,
    age INT NOT NULL
);
```

```
INSERT INTO custom_user (name, age) VALUES('Slavik', 19);
INSERT INTO custom_user (name, age) VALUES('Danil', 18);
INSERT INTO custom_user (name, age) VALUES('Dima', 20);
```
