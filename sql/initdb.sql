CREATE TABLE Photo (
    id serial,
    url text,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);



CREATE TABLE UserProfile (
     id         serial,
     email      text NOT NULL UNIQUE,
     link       text NOT NULL UNIQUE,
     avatar_id  int REFERENCES Photo(id) ON DELETE SET NULL,
     sex        text,
     bio        text,
     birthday   text,
     is_deleted boolean NOT NULL DEFAULT false,
     dying_time interval,
     PRIMARY KEY (id)
);

CREATE TABLE UserSubscriber (
    user_id int REFERENCES UserProfile(id) ON DELETE SET NULL,
    subscriber_id int REFERENCES UserProfile(id) ON DELETE SET NULL
);

CREATE Table Album (
    id serial,
    user_id int REFERENCES UserProfile(id) ON DELETE CASCADE,
    title text NOT NULL,
    visibility boolean NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE AlbumPhoto (
    album_id int REFERENCES Album(id) ON DELETE SET NULL,
    photo_id int REFERENCES Photo(id) ON DELETE CASCADE
);

CREATE TABLE Community (
    id serial,
    title text NOT NULL,
    link text NOT NULL UNIQUE,
    bio text,
    privacy text NOT NULL,
    creation_date text NOT NULL,
    id_deleted boolean NOT NULL DEFAULT false,
    is_banned boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);

CREATE TABLE CommunityManagement (
    user_id int REFERENCES UserProfile(id),
    community_id int REFERENCES Community(id),
    user_role text,
    description text
);


CREATE TABLE CommunitySubscriber (
    user_id int REFERENCES UserProfile(id),
    community_id int REFERENCES Community(id)
);

CREATE TABLE Post (
    id bigserial,
    author_id int REFERENCES UserProfile(id),
    community_id int REFERENCES Community(id),
    show_author boolean,
    text_content text,
    likes_amount int,
    creation_date text,
    change_date text,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);


CREATE TABLE Documents (
        id serial,
        url text,
        is_deleted boolean NOT NULL DEFAULT false,
        PRIMARY KEY (id)
        );

CREATE TABLE PostDocument (
        doc_id int REFERENCES Documents(id),
        post_id int REFERENCES Post(id)
);

CREATE TABLE PostPhoto (
        photo_id int REFERENCES Photo(id),
        post_id int REFERENCES Post(id)
);

CREATE TABLE PostLike (
        user_id int REFERENCES UserProfile(id),
        post_id int REFERENCES Post(id)
);

CREATE TABLE Comment (
        id           serial,
        post_id      int REFERENCES Post (id),
        user_id      int REFERENCES Post (id),
        reply_to     int REFERENCES Comment (id),
        text_content text,
        creation_date text,
        change_date  text,
        is_deleted boolean NOT NULL DEFAULT false,
        PRIMARY KEY (id)
);

CREATE TABLE CommentDocument (
        doc_id int REFERENCES Photo(id),
        comment_id int REFERENCES Comment(id)
);

CREATE TABLE CommentPhoto (
        photo_id int REFERENCES Photo(id),
        comment_id int REFERENCES Comment(id)
);

CREATE TABLE CommentLike (
        user_id int REFERENCES UserProfile(id),
        comment_id int REFERENCES Comment(id)
);

CREATE TABLE Folder (
        id serial,
        user_id integer REFERENCES UserProfile(id),
        title text,
        PRIMARY KEY (id)
);

CREATE TABLE Chat (
        id serial,
        PRIMARY KEY (id)
);

CREATE TABLE ChatFolder (
        folder_id int REFERENCES Folder(id),
        chat_id int REFERENCES Chat(id)
);


CREATE TABLE GroupChat (
        chat_id int REFERENCES Chat(id),
        avatar_id int REFERENCES Photo(id),
        title text
);

CREATE TABLE ChatMember (
        chat_id int REFERENCES Chat(id),
        user_id int REFERENCES UserProfile(id),
        role text
);


CREATE TABLE Message (
        id serial,
        user_id int REFERENCES UserProfile(id),
        chat_id int REFERENCES Chat(id),
        message_content_type text,
        text_content text,
        creation_date date,
        change_date date,
        reply_to int REFERENCES Message(id),

        is_deleted boolean NOT NULL DEFAULT false,
        PRIMARY KEY (id)
);

CREATE TABLE MessageDocument (
        doc_id int REFERENCES Documents(id),
        message_id int REFERENCES Message(id)
);

CREATE TABLE PhotoDocument (
        photo_id int REFERENCES Photo(id),
        message_id int REFERENCES Message(id)
);


CREATE TABLE StickerPack (
        id serial,
        author text,
        price int,
        date date,
        PRIMARY KEY (id)
);

CREATE TABLE Sticker (
        id serial,
        url text,
        sticker_pack_id int REFERENCES StickerPack(id),
        PRIMARY KEY (id)
);








