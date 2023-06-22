-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (нет составного превичного ключа)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE Photo
(
    id         serial,
    url        text,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (нет составного превичного ключа)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE UserProfile
(
    id              serial,
    email           text     NOT NULL UNIQUE, -- email ипользуется для входа, поэтому должен быть уникальным
    link            text UNIQUE DEFAULT '', -- линк должен быть уникальным
    first_name      text              DEFAULT '',
    last_name       text              DEFAULT '',
    password        text     NOT NULL,
    sex             text              DEFAULT '',
    bio             text              DEFAULT '',
    status          text              DEFAULT '',
    birthday        text              DEFAULT '',
    avatar_id       int      REFERENCES Photo (id) ON DELETE SET NULL,
    last_active     text              DEFAULT '',
    is_deleted      boolean  NOT NULL DEFAULT false,
    dying_time      interval NOT NULL DEFAULT INTERVAL '6 months',
    access_to_posts text     NOT NULL DEFAULT 'all',
    avg_avatar_color text default '',
    PRIMARY KEY (id)

);

-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (нет составного превичного ключа)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE Table Album
(
    id         serial,
    user_id    int REFERENCES UserProfile (id) ON DELETE CASCADE,
    title      text    NOT NULL,
    visibility boolean NOT NULL,
    PRIMARY KEY (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE AlbumPhoto
(
    album_id int REFERENCES Album (id) ON DELETE SET NULL,
    photo_id int REFERENCES Photo (id) ON DELETE CASCADE
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (нет составного превичного ключа)
-- первая нормальная форма так же соблюдается (поля атомарны)
-- ***  таблица денормализована - количество подписчиков вынесено в поле subscribers,
-- чтобы оптимизировать запросы на получение информации о сообществе, где есть поле "количество подписчиков":
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L446
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L397
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L461
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L428
CREATE TABLE groups
(
    id            serial,
    title         text    NOT NULL,
    link          text    UNIQUE, -- линк должен быть уникальным
    owner_id      int REFERENCES UserProfile(id),
    avatar_id     int REFERENCES Photo (id),
    avg_avatar_color text default '',
    group_info           text DEFAULT '',
    privacy       text    NOT NULL DEFAULT 'open' CHECK ( privacy IN ('open', 'close') ),
    creation_date text    NOT NULL,
    hide_author   boolean default false,
    is_deleted    boolean NOT NULL DEFAULT false,
    is_banned     boolean NOT NULL DEFAULT false,
    subscribers  int default 0,
    PRIMARY KEY (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE GroupManagement
(
    user_id      int REFERENCES UserProfile (id),
    group_id int REFERENCES groups (id),
    user_role    text,
    description  text
);

-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE GroupSubscriber
(
    user_id      int REFERENCES UserProfile (id),
    group_id int REFERENCES groups (id),
    accepted bool default true,
    unique (user_id, group_id) -- подписаться можно лишь однократно
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (нет составного превичного ключа)
-- первая нормальная форма так же соблюдается (поля атомарны)
-- ***  таблица денормализована - количество лайков и комментариев вынесено в поля likes_amount, comments_amount,
-- чтобы оптимизировать запросы на получение информации о посте:
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L641
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L692
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L703
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L724
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L753
CREATE TABLE Post
(
    id            bigserial,
    author_id     int REFERENCES UserProfile (id),
    group_id  int REFERENCES groups (id),
    owner_id      int REFERENCES UserProfile (id),
    show_author   boolean,
    text_content  text,
    likes_amount  int     NOT NULL DEFAULT 0,
    comments_amount int not null default 0,
    creation_date text,
    change_date   text,
    is_deleted    boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (не составной первичный ключ)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE Attachment
(
    id         serial,
    url        text,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE PostAttachment
(
    att_id  int REFERENCES Attachment (id),
    post_id int REFERENCES Post (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE PostLike
(
    user_id int REFERENCES UserProfile (id),
    post_id int REFERENCES Post (id),
    UNIQUE (user_id, post_id) -- лайкнуть можно лишь однократно
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (не составной первичный ключ)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE Comment
(
    id            serial,
    post_id       int REFERENCES Post (id),
    user_id       int REFERENCES UserProfile (id),
    reply_to      int REFERENCES Comment (id),
    text_content  text,
    creation_date text,
    change_date   text,
    is_deleted    boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE CommentDocument
(
    doc_id     int REFERENCES Photo (id),
    comment_id int REFERENCES Comment (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE CommentPhoto
(
    photo_id   int REFERENCES Photo (id),
    comment_id int REFERENCES Comment (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE CommentLike
(
    user_id    int REFERENCES UserProfile (id),
    comment_id int REFERENCES Comment (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE Folder
(
    id      serial,
    user_id integer REFERENCES UserProfile (id),
    title   text,
    PRIMARY KEY (id)
);

-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (не составной первичный ключ)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE Chat
(
    id             serial,
    members_number int NOT NULL,
    PRIMARY KEY (id)
);

-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE ChatFolder
(
    folder_id int REFERENCES Folder (id),
    chat_id   int REFERENCES Chat (id)
);

-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (не составной первичный ключ)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE GroupChat
(
    chat_id   int REFERENCES Chat (id),
    avatar_id int REFERENCES Photo (id),
    title     text
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE ChatMember
(
    chat_id int REFERENCES Chat (id) ON DELETE CASCADE,
    user_id int REFERENCES UserProfile (id),
    last_read text default '',
    role    text DEFAULT 'member' CHECK ( role in ('member', 'admin') ),
    UNIQUE (chat_id, user_id) -- участником чата можно быть лишь однократно
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (не составной первичный ключ)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE Message
(
    id                   serial,
    user_id              int REFERENCES UserProfile (id),
    chat_id              int REFERENCES Chat (id),
    message_content_type text,
    sticker_id           int,
    text_content         text,
    creation_date        text,
    change_date          text,
    reply_to             int REFERENCES Message (id) DEFAULT NULL,

    is_deleted           boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE MessageAttachment
(
    doc_id     int REFERENCES Attachment (id),
    message_id int REFERENCES Message (id)
);


-- таблица не находится в третьей нормальной форме тк есть транзитивная зависимость id -> author,  author -> depeche_authored, id -> depeche_authored
-- вторая нормальная форма соблюдается (не составной первичный ключ)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE StickerPack
(
    id     serial,
    title text,
    author text default '',
    depeche_authored bool default false,
    cover text,
    description text,
    creation_date  text,
    PRIMARY KEY (id)
);



-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых
-- вторая нормальная форма соблюдается (не составной первичный ключ)
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE Sticker
(
    id              serial,
    url             text,
    stickerpack_id int REFERENCES StickerPack (id),
    PRIMARY KEY (id)
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых (ключевые атрибуты отсутвуют)
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE userstickerpack
(
    user_id int references userprofile(id),
    pack_id int references stickerpack(id),
    unique (user_id, pack_id) -- один стикер пак можно добавить лишь однократно
);


-- таблица находится в третей нормальной форме, т.к отсутствуют транзитивные  зависимости неключевых атрибутов от ключевых (ключевые атрибуты отсутвуют)
-- вторая нормальная форма соблюдается, т.к первичные ключи отсутсвуют
-- первая нормальная форма так же соблюдается (поля атомарны)
CREATE TABLE FriendRequests
(
    subscriber   serial references userprofile (id),
    subscribed   serial references userprofile (id),
    request_time text,
    rejected     boolean default false,
    unique (subscriber, subscribed) -- можно отправить одному человеку лишь одну заявку в друзья
);

CREATE OR REPLACE FUNCTION increase_subscribers_count()
    RETURNS trigger AS
$$
BEGIN
    UPDATE groups set subscribers = subscribers + 1
    where NEW.group_id = id and NEW.accepted;
    return NEW;
END;
$$
    LANGUAGE 'plpgsql';


CREATE OR REPLACE FUNCTION decrease_subscribers_count()
    RETURNS trigger AS
$$
BEGIN
    UPDATE groups set subscribers = subscribers - 1
    where OLD.group_id = id;
    return OLD;
END;
$$
    LANGUAGE 'plpgsql';


CREATE OR REPLACE FUNCTION increase_comments_count()
    RETURNS trigger AS
$$
BEGIN
    UPDATE post set comments_amount = comments_amount + 1
    where NEW.post_id = post.id;
    return NEW;
END;
$$
    LANGUAGE 'plpgsql';


CREATE OR REPLACE FUNCTION decrease_comments_count()
    RETURNS trigger AS
$$
BEGIN
    IF new.is_deleted = false then
        return new;
    end if;
    RAISE log 'Value: %', NEW;
    UPDATE post set comments_amount = comments_amount - 1
    where new.post_id = id;
    return new;
END;
$$
    LANGUAGE 'plpgsql';



CREATE trigger increase_subscribers_count_on_insert_update_trigger
    after insert or update of accepted on groupsubscriber
    FOR EACH ROW
    EXECUTE FUNCTION increase_subscribers_count();

CREATE trigger decrease_subscribers_count_on_delete_trigger
    before delete on groupsubscriber
    FOR EACH ROW
EXECUTE FUNCTION decrease_subscribers_count();

CREATE or REPLACE trigger increase_comments_count_on_delete_trigger
    after insert on comment
    FOR EACH ROW
EXECUTE FUNCTION increase_comments_count();


CREATE or REPLACE trigger decrease_comments_count_on_delete_trigger
    after update of is_deleted on comment
    FOR EACH ROW
EXECUTE FUNCTION decrease_comments_count();

insert into stickerpack (id, title, depeche_authored, cover, description, creation_date)
values (1, 'Персик', true, '1/sticker_vk_persik_000.png', 'Величественный комок шерсти, обожающий спать, мурчать и играть с компьютерной мышкой.', ''),
       (2, 'Дигги', true, '2/sticker_vk_diggy_001.png', 'Шикарный, непревзойдённый мемасный пёс!', '');

insert into sticker (url, stickerpack_id)
values ('1/sticker_vk_persik_000.png', 1),
       ('1/sticker_vk_persik_001.png', 1),
       ('1/sticker_vk_persik_002.png', 1),
       ('1/sticker_vk_persik_003.png', 1),
       ('1/sticker_vk_persik_004.png', 1),
       ('1/sticker_vk_persik_005.png', 1),
       ('1/sticker_vk_persik_006.png', 1),
       ('1/sticker_vk_persik_007.png', 1),
       ('1/sticker_vk_persik_008.png', 1),
       ('1/sticker_vk_persik_009.png', 1),
       ('1/sticker_vk_persik_010.png', 1),
       ('1/sticker_vk_persik_011.png', 1),
       ('1/sticker_vk_persik_012.png', 1),
       ('1/sticker_vk_persik_013.png', 1),
       ('1/sticker_vk_persik_014.png', 1),
       ('1/sticker_vk_persik_015.png', 1),
       ('2/sticker_vk_diggy_001.png', 2),
       ('2/sticker_vk_diggy_002.png', 2),
       ('2/sticker_vk_diggy_003.png', 2),
       ('2/sticker_vk_diggy_004.png', 2),
       ('2/sticker_vk_diggy_005.png', 2),
       ('2/sticker_vk_diggy_006.png', 2),
       ('2/sticker_vk_diggy_007.png', 2),
       ('2/sticker_vk_diggy_008.png', 2),
       ('2/sticker_vk_diggy_009.png', 2),
       ('2/sticker_vk_diggy_010.png', 2),
       ('2/sticker_vk_diggy_011.png', 2);

insert into stickerpack (id, title, depeche_authored, cover, description, creation_date)
values (3, 'Жабеня', true, '3/sticker_zhabenya_001.webp', 'Жабы на все случаи жизни.', '');

insert into sticker(url, stickerpack_id)
values ('3/sticker_zhabenya_001.webp', 3),
       ('3/sticker_zhabenya_002.webp', 3),
       ('3/sticker_zhabenya_003.webp', 3),
       ('3/sticker_zhabenya_004.webp', 3),
       ('3/sticker_zhabenya_005.webp', 3),
       ('3/sticker_zhabenya_006.webp', 3),
       ('3/sticker_zhabenya_007.webp', 3),
       ('3/sticker_zhabenya_008.webp', 3),
       ('3/sticker_zhabenya_009.webp', 3),
       ('3/sticker_zhabenya_010.webp', 3);

insert into stickerpack (id, title, depeche_authored, cover, description, creation_date)
values (5, 'cats cats cats', true, '5/cover', 'cats, cats and also cats :з', '');


insert into sticker(url, stickerpack_id)
values ('5/sticker_cats_001.webm', 5),
       ('5/sticker_cats_002.webm', 5),
       ('5/sticker_cats_003.webm', 5),
       ('5/sticker_cats_004.webm', 5),
       ('5/sticker_cats_005.webm', 5),
       ('5/sticker_cats_006.webm', 5),
       ('5/sticker_cats_007.webm', 5),
       ('5/sticker_cats_008.webm', 5),
       ('5/sticker_cats_009.webm', 5),
       ('5/sticker_cats_010.webm', 5);


insert into stickerpack (id, title, depeche_authored, cover, description, creation_date)
values (6, 'Имбирчик', true, '6/ezh_5.webp', 'Добрый еж.', '');


insert into sticker(url, stickerpack_id)
values ('6/ezh_1.webp', 6),
       ('6/ezh_2.webp', 6),
       ('6/ezh_3.webp', 6),
       ('6/ezh_4.webp', 6),
       ('6/ezh_5.webp', 6),
       ('6/ezh_6.webp', 6),
       ('6/ezh_7.webp', 6),
       ('6/ezh_8.webp', 6),
       ('6/ezh_9.webp', 6),
       ('6/ezh_10.webp', 6);



-- индекс для запроса https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L31
CREATE INDEX IF NOT EXISTS user_link_idx
    ON UserProfile(link);

-- индекс для запроса https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L17
CREATE INDEX IF NOT EXISTS user_email_idx
    ON UserProfile(email);

-- индекс для оптимизации запроса https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L1050
CREATE INDEX IF NOT EXISTS comment_creation_date_idx
    ON Comment(creation_date);

-- индекс для оптимизации запросов
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L703
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L724
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L753
CREATE INDEX IF NOT EXISTS post_creation_date_idx
    ON Post(creation_date);

-- индекс для оптимизации запросов
-- https://github.com/go-park-mail-ru/2023_1_CashMap/blob/bc0a30922f148f4150dfa3caa6c7718bcbbcdabb/internal/repository/postgres/query.go#L777
CREATE INDEX IF NOT EXISTS message_creation_date_idx
    ON Message(creation_date);

-- foreign keys indexes --
CREATE INDEX IF NOT EXISTS group_owner_idx
    ON Groups(owner_id);

CREATE INDEX IF NOT EXISTS post_owner_idx
    ON Post(owner_id);

CREATE INDEX IF NOT EXISTS post_author_idx
    ON Post(author_id);

CREATE INDEX IF NOT EXISTS message_author_id_idx
    ON Message(user_id);

CREATE INDEX IF NOT EXISTS chat_member_user_id_idx
    ON ChatMember(user_id);

CREATE INDEX IF NOT EXISTS chat_member_chat_id_idx
    ON ChatMember(chat_id);

CREATE INDEX IF NOT EXISTS comment_user_id_idx
    ON Comment(user_id);

CREATE INDEX IF NOT EXISTS comment_post_id_idx
    ON Comment(post_id);

CREATE INDEX IF NOT EXISTS userprofile_avatar_id_idx
    ON UserProfile(avatar_id);
-- foreign keys indexes --


SELECT pg_size_pretty(sum(pg_relation_size(idx))::bigint) as size,
       (array_agg(idx))[1] as idx1, (array_agg(idx))[2] as idx2,
       (array_agg(idx))[3] as idx3, (array_agg(idx))[4] as idx4
FROM (
         SELECT indexrelid::regclass as idx, (indrelid::text ||E'\n'|| indclass::text ||E'\n'|| indkey::text ||E'\n'||
                                              coalesce(indexprs::text,'')||E'\n' || coalesce(indpred::text,'')) as key
         FROM pg_index) sub
GROUP BY key HAVING count(*)>1
ORDER BY sum(pg_relation_size(idx)) DESC;