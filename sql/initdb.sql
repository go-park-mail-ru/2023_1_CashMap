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
    link          text    NOT NULL UNIQUE,
    avatar_id     int REFERENCES Photo (id),
    group_info           text DEFAULT '',
    privacy       text    NOT NULL DEFAULT 'open' CHECK ( privacy IN ('open', 'default', 'close') ),
    creation_date text    NOT NULL,
    hide_author   boolean default false,
    id_deleted    boolean NOT NULL DEFAULT false,
    is_banned     boolean NOT NULL DEFAULT false,
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
    group_id int REFERENCES groups (id)
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
    creation_date text,
    change_date   text,
    is_deleted    boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);


CREATE TABLE Documents
(
    id         serial,
    url        text,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);

CREATE TABLE PostDocument
(
    doc_id  int REFERENCES Documents (id),
    post_id int REFERENCES Post (id)
);

CREATE TABLE PostPhoto
(
    photo_id int REFERENCES Photo (id),
    post_id  int REFERENCES Post (id)
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
    user_id       int REFERENCES Post (id),
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
    reply_to             int REFERENCES Message (id),

    is_deleted           boolean NOT NULL DEFAULT false,
    PRIMARY KEY (id)
);

CREATE TABLE MessageDocument
(
    doc_id     int REFERENCES Documents (id),
    message_id int REFERENCES Message (id)
);

CREATE TABLE PhotoDocument
(
    photo_id   int REFERENCES Photo (id),
    message_id int REFERENCES Message (id)
);


CREATE TABLE StickerPack
(
    id     serial,
    author text,
    price  int,
    date   date,
    PRIMARY KEY (id)
);

CREATE TABLE Sticker
(
    id              serial,
    url             text,
    sticker_pack_id int REFERENCES StickerPack (id),
    PRIMARY KEY (id)
);


    CREATE TABLE FriendRequests
(
    subscriber   serial references userprofile (id),
    subscribed   serial references userprofile (id),
    request_time text,
    rejected     boolean default false,
    unique (subscriber, subscribed)
)





