CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL,
    first_name VARCHAR(150),
    last_name VARCHAR(150),
    username VARCHAR(150) ,
    password varchar(256) ,
    email VARCHAR(150) ,
    picture VARCHAR(256) ,
    created_at timestamp DEFAULT now(),
    updated_at timestamp ,
    CONSTRAINT pk_users PRIMARY KEY(id)
);


CREATE TABLE IF NOT EXISTS articles (
    id serial NOT NULL,
    name VARCHAR(150) ,
    price INT default 0,
    picture VARCHAR(256) , 
    description text ,
    user_id int ,
    created_at timestamp DEFAULT now(),
    updated_at timestamp ,
    CONSTRAINT pk_notes PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS pedido (
    id serial NOT NULL,
    estado VARCHAR(150),	
    user_id int NOT NULL,
    article_id int NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp ,
    CONSTRAINT pk_notes PRIMARY KEY(id),
    CONSTRAINT fk_pedido_users FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_article_users FOREIGN KEY(article_id) REFERENCES articles(id)
    	
);

INSERT INTO users (first_name, last_name, username, email) VALUES ('Juan', 'Perez', 'juanperez', 'juanperez@email.com');
INSERT INTO users (first_name, last_name, username, email) VALUES ('Jhon', 'Doe', 'jhondoe', 'jhondoe@email.com');
INSERT INTO users (first_name, last_name, username, email) VALUES ('Pedro', 'Pereira', 'pedropereira', 'pedropereira@email.com');
INSERT INTO users (first_name, last_name, username, email) VALUES ('Maria', 'Castro', 'mariacastro', 'mariacastro@email.com');
INSERT INTO users (first_name, last_name, username, email) VALUES ('Cristina', 'Colorado', 'cristinacolorado', 'cristinacolorado@email.com');


INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('Tetris', 200000, 'https://i.blogs.es/91fe1d/180719-mejores1989-01/1366_2000.jpeg', 'Empezamos por ese vendeconsolas (literalmente) llamado Tetris, más en concreto su adaptación a la Game Boy de Nintendo que popularizó hasta el extremo la fiebre por los videojuegos y que quedase patente que no todo eran disparos en este mundillo. Un clásico de 1984 que no deja de renovarse.', 1);
INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('Teenage Mutant Ninja Turtles', 150000, 'https://i.blogs.es/05eaa8/180719-mejores1989-02/1366_2000.jpeg', 'Recientemente recordamos la mítica recreativa de las Tortugas Ninja, obra de Konami. Teenage Mutant Ninja Turtles fue un beat em up que entraba por la vista y que era disfrutable a lo grande al juntarse con tres personas más. Logró, junto con el plataformas de acción homónimo para NES (de 1989, también), que el interés por la serie animada (de 1987) aumentase.',2);
INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('Super Off Road', 170000, 'https://i.blogs.es/1737b5/180719-mejores1989-03/1366_2000.jpeg', 'Bajo el nombre del legendario Ivan Ironman Stewart como principal reclamo comercial para dar a conocer este Super Off Road, el estudio Leland Corporation supo crear un juego de carreras arcade de lo más directo y con identidad propia, fácilmente reconocible pese al paso de tantos años. El pique entre colegas podía llegar a ser considerable con sus colisiones.',2);
INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('DuckTales', 210000, 'https://i.blogs.es/b893b2/180719-mejores1989-04/1366_2000.jpeg', 'Un plataformas muy recordado de la etapa de NES es el adorable DuckTales de Capcom, tanto por controlar al tío Gilito haciendo pogo con su bastón, como por la fantástica banda sonora. Su remake para PC, PS3, Wii U y Xbox 360 no gozó del mismo encanto pese a contar con un apartado gráfico muy atractivo. Al menos se rescató el original en The Disney Afternoon Collection.', 3);
INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('Populous', 200000, 'https://i.blogs.es/8b6f9e/180719-mejores1989-05/1366_2000.jpeg', 'Populous, el simulador de Dios por excelencia. El que lo inició todo. El Big-Bang de Peter Molyneux y una de las obras más reconocibles del mítico estudio Bullfrog Productions. Un juego de estrategia diferente, casi divino, donde optamos por hacer el bien... y también a ahogar a sus habitantes.', 1);
INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('SimCity', 300000, 'https://i.blogs.es/eb7fb5/180719-mejores1989-06/1366_2000.jpeg', 'Otro clásico de la estrategia, en este caso de gestión de ciudades. A SimCity le debemos, además, la posterior experimentación con Los Sims en el nuevo milenio, como clara evolución de aquel concepto primigenio de Will Wright. Aquella obra de 1989 apostó por la vista cenital, que estaba muy de moda.', 4);
INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('Golden Axe', 250000, 'https://i.blogs.es/13995c/180719-mejores1989-07/1366_2000.jpeg', 'SEGA perfeccionó su fórmula con los beat em up iniciando una nueva senda de gloria con la saga Golden Axe, hasta que llegó a su perfección absoluta en aquel Golden Axe: Revenge of Death Adder de recreativas en 1992. La obra original, eso sí, se mantiene fresca y sale en las colecciones de SEGA.', 5);
INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('Shadow of the Beast', 415000, 'https://i.blogs.es/a6d7ab/180719-mejores1989-08/1366_2000.jpeg', 'Shadow of the Beast es esa clase de juegos que asocias irremediablemente con el Amiga de Commodore. Una aventura que supo aprovechar casi todo su potencial para deslumbrar al personal. Eso sí, era duro como pocos y formó parte de una trilogía memorable por parte de Psygnosis. Ahora bien, conviene olvidar aquel intento por resurgir su marca como exclusivo en PS4.', 2);
INSERT INTO ARTICLES (name, price, picture,  description, user_id)
        VALUES ('Pipe Mania', 150000, 'https://i.blogs.es/b615ec/180719-mejores1989-09/1366_2000.jpeg', 'Seguro que todo aquel que vea la imagen de arriba recordará al instante el mítico Pipe Mania del estudio The Assembly Line. Porque su mecánica fue toda una novedad y ha sido copiada hasta la saciedad desde entonces, hasta el punto de formar parte de muchos minijuegos, como en el gran BioShock.', 1);
INSERT INTO ARTICLES (name, price, picture, description, user_id)
        VALUES ('Batman', 100000, 'https://i.blogs.es/ac994a/180719-mejores1989-10/1366_2000.jpeg', 'De los múltiples videojuegos oficiales de Batman, basados en la película homónima de Tim Burton, destacó especialmente aquel de NES al que se le sigue recordando por su legendaria dificultad en las últimas fases. Todo un clásico de Sunsoft que nos encantaría poder jugar algún día en Switch.', 2);

INSERT INTO pedido (estado,user_id,article_id) VALUES ('entregado',2,3);
INSERT INTO pedido (estado,user_id,article_id) VALUES ('en despacho',1,2);
INSERT INTO pedido (estado,user_id,article_id) VALUES ('entregado',4,1);        