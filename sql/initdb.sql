CREATE TABLE Photo
(
    id         serial,
    url        text,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);


CREATE TABLE UserProfile
(
    id              serial,
    email           text     NOT NULL UNIQUE,
    link            text UNIQUE DEFAULT '',
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
    PRIMARY KEY (id)

);


CREATE Table Album
(
    id         serial,
    user_id    int REFERENCES UserProfile (id) ON DELETE CASCADE,
    title      text    NOT NULL,
    visibility boolean NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE AlbumPhoto
(
    album_id int REFERENCES Album (id) ON DELETE SET NULL,
    photo_id int REFERENCES Photo (id) ON DELETE CASCADE
);

CREATE TABLE groups
(
    id            serial,
    title         text    NOT NULL,
    link          text    UNIQUE,
    owner_id      int REFERENCES UserProfile(id),
    avatar_id     int REFERENCES Photo (id),
    group_info           text DEFAULT '',
    privacy       text    NOT NULL DEFAULT 'open' CHECK ( privacy IN ('open', 'close') ),
    creation_date text    NOT NULL,
    hide_author   boolean default false,
    is_deleted    boolean NOT NULL DEFAULT false,
    is_banned     boolean NOT NULL DEFAULT false,
    subscribers  int default 0,
    PRIMARY KEY (id)
);

CREATE TABLE GroupManagement
(
    user_id      int REFERENCES UserProfile (id),
    group_id int REFERENCES groups (id),
    user_role    text,
    description  text
);


CREATE TABLE GroupSubscriber
(
    user_id      int REFERENCES UserProfile (id),
    group_id int REFERENCES groups (id),
    accepted bool default true,
    unique (user_id, group_id)
);

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


CREATE TABLE Attachment
(
    id         serial,
    url        text,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);

CREATE TABLE PostAttachment
(
    att_id  int REFERENCES Attachment (id),
    post_id int REFERENCES Post (id)
);



CREATE TABLE PostLike
(
    user_id int REFERENCES UserProfile (id),
    post_id int REFERENCES Post (id),
    UNIQUE (user_id, post_id)
);

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

CREATE TABLE CommentDocument
(
    doc_id     int REFERENCES Photo (id),
    comment_id int REFERENCES Comment (id)
);

CREATE TABLE CommentPhoto
(
    photo_id   int REFERENCES Photo (id),
    comment_id int REFERENCES Comment (id)
);

CREATE TABLE CommentLike
(
    user_id    int REFERENCES UserProfile (id),
    comment_id int REFERENCES Comment (id)
);

CREATE TABLE Folder
(
    id      serial,
    user_id integer REFERENCES UserProfile (id),
    title   text,
    PRIMARY KEY (id)
);

CREATE TABLE Chat
(
    id             serial,
    members_number int NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE ChatFolder
(
    folder_id int REFERENCES Folder (id),
    chat_id   int REFERENCES Chat (id)
);


CREATE TABLE GroupChat
(
    chat_id   int REFERENCES Chat (id),
    avatar_id int REFERENCES Photo (id),
    title     text
);

CREATE TABLE ChatMember
(
    chat_id int REFERENCES Chat (id) ON DELETE CASCADE,
    user_id int REFERENCES UserProfile (id),
    role    text DEFAULT 'member' CHECK ( role in ('member', 'admin') ),
    UNIQUE (chat_id, user_id)
);


CREATE TABLE Message
(
    id                   serial,
    user_id              int REFERENCES UserProfile (id),
    chat_id              int REFERENCES Chat (id),
    message_content_type text,
    text_content         text,
    creation_date        text,
    change_date          text,
    reply_to             int REFERENCES Message (id) DEFAULT NULL,

    is_deleted           boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);

CREATE TABLE MessageAttachment
(
    doc_id     int REFERENCES Attachment (id),
    message_id int REFERENCES Message (id)
);


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

CREATE TABLE Sticker
(
    id              serial,
    url             text,
    stickerpack_id int REFERENCES StickerPack (id),
    PRIMARY KEY (id)
);

CREATE TABLE userstickerpack
(
    user_id int references userprofile(id),
    pack_id int references stickerpack(id),
    unique (user_id, pack_id)
);


    CREATE TABLE FriendRequests
(
    subscriber   serial references userprofile (id),
    subscribed   serial references userprofile (id),
    request_time text,
    rejected     boolean default false,
    unique (subscriber, subscribed)
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
    where OLD.group_id = id and OLD.accepted;
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