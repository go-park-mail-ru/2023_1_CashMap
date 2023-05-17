package postgres

// TODO sink fields with entities
var (
	UserById = `
	select  u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
		case when p.url is null 
			then ''
			else p.url end avatar
	from userprofile u
         left join photo p on u.avatar_id = p.id
	where u.id = $1
	`
	UserByEmail = `
	select u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
		u.password,
       case when p.url is null 
           then ''
           else p.url end avatar
	from userprofile u
         left join photo p on u.avatar_id = p.id
	where u.email = $1
	`

	UserByLink = `
	select u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
       case when p.url is null 
           then ''
           else p.url end avatar
	from userprofile u
         left join photo p on u.avatar_id = p.id
	where u.link = $1
	`

	FriendsById = `
	select u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
	  case when p.url is null
	      then '' 
	      else p.url end avatar from
	friendrequests f1
	join friendrequests f2 on
	   f1.subscribed = f2.subscriber and
	   f2.subscribed = f1.subscriber
	join userprofile u on
	   f1.subscribed = u.id
	left join photo p on
	  u.avatar_id = p.id
	where
	   f1.subscriber = $1
	and u.id > $2
	limit $3
	`

	SubscribesById = `
	select u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
		case when p.url is null
            then ''
        	else p.url
    	end avatar
	from
    friendrequests f1
        join friendrequests f2
            on f1.subscriber = f2.subscribed and 
               f2.subscriber = f1.subscribed
        right join friendrequests f3
            on f3.subscriber = f1.subscriber and 
               f1.subscribed = f3.subscribed
        join userprofile u
            on f3.subscribed = u.id
        left join photo p
            on u.avatar_id = p.id
	where
    	f1 is null and f3.subscriber = $1
	and u.id > $2
	limit $3
	`

	SubscribersById = `
	select u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
		case when p.url is null
            then ''
        	else p.url
    	end avatar
	from
    friendrequests f1
        join friendrequests f2
            on f1.subscriber = f2.subscribed and 
               f2.subscriber = f1.subscribed
        right join friendrequests f3
            on f3.subscriber = f1.subscriber and 
               f1.subscribed = f3.subscribed
        join userprofile u
            on f3.subscriber = u.id
        left join photo p
            on u.avatar_id = p.id
	where
    	f1 is null and f3.subscribed = $1
	and u.id > $2
	order by u.id
	limit $3
	`

	PendingFriendRequestsById = `
	select u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
		case when p.url is null
                then ''
            else p.url
           end avatar
	from
    friendrequests f1
        join friendrequests f2
             on f1.subscriber = f2.subscribed and
                f2.subscriber = f1.subscribed
        right join friendrequests f3
                   on f3.subscriber = f1.subscriber and
                      f1.subscribed = f3.subscribed
        join userprofile u
             on f3.subscriber = u.id
        left join photo p
                  on u.avatar_id = p.id
	where
    f1 is null and f3.subscribed = $1 and f3.rejected = false
	and u.id > $2
	order by u.id
	limit $3
	`

	Subscribe = `
	with sub as
         (insert into friendrequests
             (subscriber, subscribed, request_time)
             select u1.id , u2.id, $3 from
                userprofile u1 cross join
                userprofile u2 where u1.email = $1 and u2.link = $2
             returning
                 subscriber s1, subscribed s2, request_time time
         )
	update friendrequests
	set rejected = false,
    	request_time = time
	from sub s	
	where
        s.s1 = subscribed and
        s.s2 = subscriber 
	`
	Unsubscribe = `
	with unsub as (
    delete from friendrequests
        where (subscriber, subscribed)
        in
        (select u1.id , u2.id from
                userprofile u1 cross join
                userprofile u2 where
                u1.email = $1 and u2.link = $2)
        returning
            subscriber s1, subscribed s2
	)
	update friendrequests
	set rejected = true
	from unsub s
	where
        s.s1 = subscribed and
        s.s2 = subscriber
`

	RejectFriendRequest = `
	update friendrequests f1
	set rejected = true
	from
    	friendrequests f2
    join
        userprofile u1
            on f2.subscribed = u1.id
    join userprofile u2
            on f2.subscriber = u2.id
    where
        u1.email = $1 and u2.link = $2 and
        f1.subscriber = f2.subscriber and f1.subscribed = f2.subscribed;
`

	CreateUser = `
	insert into 
    userprofile 
    (email, password, first_name, last_name, last_active) 
	values 
    ($1, $2, $3, $4, $5) returning id
	 
`

	UpdateUserLink = `
	update userprofile 
	set 
	    link = $1
	where 
		id = $2
	`

	DeleteUser = `
	update userprofile 
	set is_deleted = true
	where email = $1
    returning id
    `

	RandomUsers = `
	with subs as (
    	select f.subscriber, f.subscribed
    		from userprofile u
        join friendrequests f
            on u.id = f.subscribed or u.id = f.subscriber
    	where u.email = $1
	)
	select
       u.id, u.link, u.email,
       u.first_name, u.last_name,
       u.sex, u.bio, u.status,
       u.birthday, u.last_active,
       case when
           p.url is null
       then
           ''
       else
           p.url
       end avatar
	from userprofile u
		left join photo p
    		on u.avatar_id = p.id
		left join subs s
			on u.id = s.subscriber or u.id = s.subscribed
	where
    	s is null and
    	u.id > $2
		and u.email <> $1
	order by u.id
	limit $3;
`

	CheckLink = `
	select exists(select * from userprofile where link = $1 ) ex
	`

	UpdateAvatar = `
	with inserted as (
    insert into photo (url) values ($1)
           returning id p_id
	)
	update userprofile u1
	set avatar_id = av.p_id
	from inserted av where u1.email = $2
	`
	// IsFriend returns true when $1 is subscribed on $2 and vice versa
	IsFriend = `
	select exists(select * from friendrequests f1
    join friendrequests f2 on
            f1.subscribed = f2.subscriber and
            f2.subscribed = f1.subscriber
where f1.subscriber = $1 and f1.subscribed = $2)
	`

	// IsSubscriber returns true when $1 is subscribed on $2
	IsSubscriber = `
select exists(select * from friendrequests f1
    left join friendrequests f2 on
            f1.subscribed = f2.subscriber and
            f2.subscribed = f1.subscriber
where f1.subscriber = $1 and f1.subscribed = $2 and 
      f2 is null)`

	// IsSubscribed returns true when $2 is subscribed on $1 (rejected request)
	IsSubscribed = `
	select exists(select * from friendrequests f1
    left join friendrequests f2 on
            f1.subscribed = f2.subscriber and
            f2.subscribed = f1.subscriber
	where f1.subscriber = $2 and f1.subscribed = $1 and 
	      f2 is null and
	      f1.rejected)
`
	// HasPendingRequest returns true when $2 is subscribed on $1 (unseen yet request)
	HasPendingRequest = `
	select exists(select * from friendrequests f1
    left join friendrequests f2 on
            f1.subscribed = f2.subscriber and
            f2.subscribed = f1.subscriber
	where f1.subscriber = $2 and f1.subscribed = $1 and 
	      f2 is null and
	      not f1.rejected)
	`
)

var (
	CreateMessage = `
	insert into message
	(user_id, chat_id, message_content_type, text_content, creation_date, reply_to)
	values 
	($1,$2,$3,$4,$5,$6)
	returning (id)
	`

	MessageById = `
	select id, user_id, 
	           chat_id, message_content_type, 
	           text_content, creation_date, 
	           reply_to, is_deleted
	from message where id = $1
`

	GetMembersByChatId = `
	select u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active,
		case when p.url is null
                then ''
            else p.url
	end avatar
	from chatmember cm join 
	    userprofile u on cm.user_id = u.id
	left join photo p on u.avatar_id = p.id
	where cm.chat_id = $1
	`
)

var (
	GroupByLink = `
	select g.title, g.link, g.group_info info, g.privacy,
	       g.creation_date, g.hide_author, u.link owner_link,
	       g.is_deleted, g.subscribers,
	case when p.url is null 
           then ''
           else p.url end avatar
	from groups g
         left join photo p on g.avatar_id = p.id
		 join userprofile u on g.owner_id = u.id
	where g.link = $1
	`

	GroupsByUserlink = `
	select g.title, g.link, g.group_info info, g.privacy,
	       g.creation_date, g.hide_author, u2.link owner_link,
	       g.is_deleted, g.subscribers,
	case when p.url is null 
           then ''
           else p.url end avatar
	from groups g 
		left join photo p on g.avatar_id = p.id
	join groupsubscriber g2 on g.id = g2.group_id and g2.accepted
	join userprofile u on g2.user_id = u.id
	join userprofile u2 on g.owner_id = u2.id
	where u.link = $1 and g.is_deleted = false
	and g.id > $3 
	order by g.id
	limit $2
	`

	GroupsByUserEmail = `
	select g.title, g.link, g.group_info info, g.privacy,
	       g.creation_date, g.hide_author, u2.link owner_link,
	       g.is_deleted, g.subscribers,
	case when p.url is null 
           then ''
           else p.url end avatar
	from groups g 
		left join photo p on g.avatar_id = p.id
	join groupsubscriber g2 on g.id = g2.group_id and g2.accepted
	join userprofile u on g2.user_id = u.id
	join userprofile u2 on g.owner_id = u2.id
	where u.email = $1 and g.is_deleted = false
	and g.id > $3 
	order by g.id
	limit $2
	`
	// TODO Сделать таблицу для обновления по крон таске или процедуре
	GetGroups = `
		select g.title, g.link, g.group_info info, g.privacy,
	       g.creation_date, g.hide_author, u.link owner_link,
	       g.is_deleted, g.subscribers,
	case when p.url is null 
           then ''
           else p.url end avatar
	from groups g 
		left join photo p on g.avatar_id = p.id
		join userprofile u on g.owner_id = u.id
		where g.is_deleted = false
	order by g.subscribers
	limit $1 offset $2
	`

	GetManaged = `
	select g.title, g.link, g.group_info info, g.privacy,
	       g.creation_date, g.hide_author, u.link owner_link,
	       g.is_deleted, g.subscribers,
	case when p.url is null 
           then ''
           else p.url end avatar
	from groups g 
		left join photo p on g.avatar_id = p.id
		join userprofile u on g.owner_id = u.id
		where u.email = $1 and g.is_deleted = false
		and g.id > $3
		order by g.id
		limit $2
	`

	GroupSubscribers = `
	select  u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
		case when p.url is null 
			then ''
			else p.url end avatar
	from userprofile u
         left join photo p on u.avatar_id = p.id
	join groupsubscriber g on u.id = g.user_id and g.accepted 
	join groups gr on g.group_id = gr.id
	where gr.link = $1
	and u.id > $3
	order by u.id
	limit $2
	`

	PendingGroupRequests = `
	select  u.id, u.link, u.email,
	    u.first_name, u.last_name, 
		u.sex, u.bio, u.status,
		u.birthday, u.last_active, 
		case when p.url is null 
			then ''
			else p.url end avatar
	from userprofile u
         left join photo p on u.avatar_id = p.id
	join groupsubscriber g on u.id = g.user_id and not g.accepted
	where u.id = $1
	and u.id > $3
	order by u.id
	limit $2`

	AcceptRequest = `
	update groupsubscriber g
	set accepted = true
	from groupsubscriber g1
    	join userprofile u on g1.user_id = u.id
    	join groups g2 on g1.group_id = g2.id
	where u.link = $1 and g2.link = $2
  	and g.user_id = u.id and g.group_id = g2.id
	`

	DeclineRequest = `
	delete from groupsubscriber g 
	where  g.user_id in (select id from userprofile u where u.link = $1)
	and g.group_id in (select id from groups g2 where g2.link = $2)
	`

	AcceptAllRequests = `
	update groupsubscriber g
	set accepted = true
	from groupsubscriber g1
    	join groups g2 on g1.group_id = g2.id
	where g2.link = $1 
  	and g.group_id = g2.id
  	`

	CreateGroup = `
		with owner as (select id from userprofile where email = $1)
		insert into groups (title, group_info, privacy, creation_date, hide_author, owner_id)
		select $2, $3, $4, $5, $6, owner.id from owner
		returning id`

	UpdateGroupLink = `
		update groups 
			set link = $1
		where id = $2`

	UpdateGroupAvatar = `
	with inserted as (
    insert into photo (url) values ($1)
        returning id p_id
	)
	update groups p1
	set avatar_id = av.p_id
	from inserted av where p1.id = $2
	`

	UpdateAvatarGroupByLink = `
	with inserted as (
    insert into photo (url) values ($1)
        returning id p_id
	)
	update groups p1
	set avatar_id = av.p_id
	from inserted av where p1.link = $2
		
`
	DeleteGroup = `
	update groups 
		set is_deleted = true
	where link = $1`

	GroupSubscribe = `
		insert into groupsubscriber (user_id, group_id, accepted) 
		select u.id, g.id, 
		       case when g.privacy = 'open'
			   then true
			   else false end accepted
		from userprofile u cross join groups g 
		where u.email = $1 and g.link = $2
	`

	GroupUnsubscribe = `
		delete from groupsubscriber g 
		where g.group_id = (select id from groups g2 where g2.link = $2 )
		and g.user_id = (select id from userprofile u where u.email = $1)
	`

	IsOwner = `
	select case when 
	    exists(select u.id 
	           from userprofile u 
			   join groups g on u.id = g.owner_id 
	           where u.email = $1 and g.link = $2)
		then true
		else false end is_owner
	`
	CheckSub = `
		select case when 
	    exists(select u.id 
	           from userprofile u 
			   join groupsubscriber g on u.id = g.user_id
	           join groups g2 on g.group_id = g2.id
	           where u.email = $1 and g2.link = $2)
		then true
		else false end is_sub
	`
	// TODO
	CheckAdminQuery = `
		select case when 
	    exists(select u.id 
	           from userprofile u 
			   join groups g on u.id = g.owner_id
	           where u.email = $1 and g.link = $2)
		then true
		else false end is_sub
	`

	InsertNewAdminQuery = `
		INSERT INTO GroupManagement (user_id, group_id, user_role)
		VALUES ((SELECT id FROM UserProfile WHERE email = $1), $2, $3)
	`
)

var (
	SetLikeQuery = `
	INSERT INTO Postlike (post_id, user_id)
	VALUES ($1, (SELECT id FROM UserProfile WHERE email = $2))
	`

	CancelLikeQuery = `
	DELETE FROM PostLike
	WHERE post_id = $1
	  AND user_id = (SELECT id FROM UserProfile WHERE email = $2)
	`

	FriendsPostsQuery = `
		SELECT  post.id,
		        text_content,
		        author.link as author_link,
		        post.likes_amount, 
        		case when post.show_author is null then true else post.show_author end  AS show_author,
        		post.creation_date,
        		CASE WHEN like_table.post_id is null THEN FALSE ELSE TRUE END as is_liked,
        		comments_amount
		FROM Post post
		    JOIN userprofile author on author.id = post.author_id 
			LEFT JOIN PostLike as like_table ON like_table.user_id = (SELECT id FROM UserProfile WHERE email = $1) AND like_table.post_id = post.id
		
		WHERE owner_id IN (SELECT u.id 
		                   FROM friendrequests f1
		                       
        JOIN friendrequests f2 on
                f1.subscribed = f2.subscriber and
                f2.subscribed = f1.subscriber
		                       
		JOIN userprofile u on
			f1.subscribed = u.id        
		WHERE f1.subscriber = (SELECT id FROM UserProfile WHERE email = $1))
		  and creation_date > $3 AND NOT post.is_deleted
		
        ORDER BY creation_date
        LIMIT $2
	`
)

var (
	PostSenderInfoQuery = `
		SELECT first_name, last_name, url, link
		FROM Post as post
				 JOIN UserProfile as profile ON post.author_id = profile.id
				 LEFT JOIN Photo as photo ON profile.avatar_id = photo.id
		WHERE post.id = $1
		`

	CommunityPostInfoQuery = `
		SELECT community.title, community.link,
		case when p.url is null 
           then ''
           else p.url end url
		FROM Post as post
				 JOIN groups as community
					  ON post.group_id = community.id
				  left join photo p on community.avatar_id = p.id
		WHERE post.id = $1
		`

	PostInfoByIdQuery = `
			SELECT post.id, text_content, author.link as author_link, post.likes_amount, post.show_author, post.creation_date, comments_amount, CASE WHEN like_table.post_id is null THEN FALSE ELSE TRUE END as is_liked
			FROM Post AS post
					 JOIN UserProfile AS author ON post.author_id = author.id
					 LEFT JOIN groups as community on post.group_id = community.id
					 LEFT JOIN UserProfile as owner ON post.owner_id = owner.id
					 LEFT JOIN PostLike as like_table ON like_table.user_id = (SELECT id FROM UserProfile WHERE email = $2) AND like_table.post_id = post.id
			WHERE post.id = $1
			  AND post.is_deleted = false
	`

	PostByCommunityLinkQuery = `
		SELECT post.id,
			   text_content,
			   author.link as author_link,
			   post.likes_amount,
			   post.show_author,
			   post.creation_date,
			   post.change_date,
			   CASE WHEN like_table.post_id is null THEN FALSE ELSE TRUE END as is_liked,
			   comments_amount
		FROM Post AS post
				 JOIN UserProfile AS author ON post.author_id = author.id
				 LEFT JOIN groups as community on post.group_id = community.id
				 LEFT JOIN PostLike as like_table ON like_table.user_id = (SELECT id FROM UserProfile WHERE email = $4) AND like_table.post_id = post.id
		WHERE post.group_id = (SELECT id FROM groups WHERE link = $1)
		  AND post.creation_date > $2
		  AND post.is_deleted = false
		ORDER BY post.creation_date DESC
		LIMIT $3
	`

	PostsByUserLinkQuery = `
		SELECT post.id,
			   text_content,
			   author.link as author_link,
			   post.likes_amount,
			   post.show_author,
			   post.creation_date,
			   post.change_date,
			   CASE WHEN like_table.post_id is null THEN FALSE ELSE TRUE END as is_liked,
			   post.comments_amount
		FROM Post AS post
				 JOIN UserProfile AS author ON post.author_id = author.id
				 LEFT JOIN groups as community on post.group_id = community.id
				 LEFT JOIN UserProfile as owner ON post.owner_id = owner.id
				LEFT JOIN PostLike as like_table ON like_table.user_id = (SELECT id FROM UserProfile WHERE email = $4) AND like_table.post_id = post.id
		WHERE post.owner_id = (SELECT id FROM UserProfile WHERE link = $1)
		  AND post.creation_date > $2
		  AND post.is_deleted = false
		ORDER BY post.creation_date DESC
		LIMIT $3
	`

	CreatePostQuery = `
		INSERT INTO Post (group_id, author_id, owner_id, show_author, text_content, creation_date, change_date)
		VALUES (:community_id, (SELECT id FROM UserProfile WHERE email = :sender_email), :owner_id, :show_author, :text,
				:init_time, :change_time)
		RETURNING id
	`
)

var (
	MessageByChatIdQuery = `
		SELECT msg.id, msg.chat_id, text_content, msg.creation_date, msg.change_date, msg.reply_to, msg.is_deleted
		FROM Message AS msg
				 JOIN UserProfile AS author ON msg.user_id = author.id
		WHERE msg.chat_id = (SELECT id FROM Chat WHERE id = $1)
		  AND msg.creation_date > $2
		  AND msg.is_deleted = false
		ORDER BY msg.creation_date DESC
		LIMIT $3
	`

	ChatsQuery = `
		SELECT chat.id as chat_id
		FROM ChatMember as member
				 JOIN Chat ON chat_id = chat.id
		WHERE member.user_id = (SELECT id FROM UserProfile WHERE email = $1 LIMIT $2 OFFSET $3)
	`

	UserInfoByChatIdQuery = `
		SELECT first_name, last_name, url, link
		FROM ChatMember
				 JOIN UserProfile ON id = user_id
				 LEFT JOIN Photo AS photo ON avatar_id = photo.id
		WHERE chat_id = $1
	`

	UserInfoByMessageIdQuery = `
		SELECT first_name, last_name, url, link
		FROM Message as msg
				 JOIN UserProfile as profile ON msg.user_id = profile.id
				 LEFT JOIN Photo as photo ON profile.avatar_id = photo.id
		WHERE msg.id = $1
	`

	IsChatExistsQuery = `
		SELECT f.chat_id
		FROM chatmember AS f
				 JOIN chatmember AS S on f.chat_id = s.chat_id
				 JOIN chat c on c.id = f.chat_id
		WHERE f.user_id = (SELECT id FROM userprofile WHERE link = $1)
		  and s.user_id = (SELECT id FROM userprofile WHERE email = $2)
		  AND f.user_id != s.user_id
		  AND c.members_number = 2
	`

	IsPersonalChatExistsQuery = `
		SELECT f.chat_id
       FROM chatmember AS f
                JOIN chatmember AS S on f.chat_id = s.chat_id
                JOIN chat c on c.id = f.chat_id
       WHERE s.user_id = (SELECT id FROM userprofile WHERE link = $1)
       group by f.chat_id
       having COUNT(*) = 1
	`

	IsChatMemberQuery = `
		SELECT true
		FROM ChatMember member
				 JOIN Chat as chat on chat.id = member.chat_id
		WHERE chat_id = $1
		  AND member.user_id = (SELECT id FROM UserProfile WHERE email = $2)
	`
)

var (
	GetUserInfoForSearchQuery = `
		WITH isFriend AS (
			select exists(select * from friendrequests f1
											join friendrequests f2 on
						f1.subscribed = f2.subscriber and
						f2.subscribed = f1.subscriber
						  where f1.subscriber = $2 and f1.subscribed = $1) as is_friend
		), isSubscriber AS (
			select exists(select * from friendrequests f1
											left join friendrequests f2 on
						f1.subscribed = f2.subscriber and
						f2.subscribed = f1.subscriber
						  where f1.subscriber = $1 and f1.subscribed = $1 and
							  f2 is null) as is_subscriber
		), isSubscribed AS (
			select exists(select * from friendrequests f1
											left join friendrequests f2 on
						f1.subscribed = f2.subscriber and
						f2.subscribed = f1.subscriber
						  where f1.subscriber = $1 and f1.subscribed = $2 and
							  f2 is null and
							  not f1.rejected) as is_subscribed
		)
		SELECT first_name, last_name, link, url, is_friend, is_subscriber, is_subscribed
		FROM UserProfile AS profile
				 LEFT JOIN Photo ON profile.avatar_id = photo.id
				 CROSS JOIN isFriend
				 CROSS JOIN isSubscriber
				 CROSS JOIN isSubscribed
		WHERE profile.id = $2;
	`

	GetCommunityInfoForSearchQuery = `
		WITH isSubscribed as (SELECT CASE WHEN COUNT(*) = 0 THEN FALSE ELSE TRUE END as is_subscribed
						  FROM GroupSubscriber WHERE user_id = (SELECT id FROM userprofile WHERE email = $2) AND group_id = $1)
		SELECT title, url, link, is_subscribed
		FROM groups
			 LEFT JOIN Photo ON groups.avatar_id = photo.id
			 CROSS JOIN isSubscribed
		WHERE groups.id = $1
	`
)

var (
	GetFeedQuery = `
	SELECT post.id,
       text_content,
       author.link as author_link,
       post.likes_amount,
       post.show_author,
       post.creation_date,
       post.change_date,
	   post.comments_amount,
       CASE WHEN like_table.post_id is null THEN FALSE ELSE TRUE END as is_liked
		FROM Post AS post
         JOIN UserProfile AS author ON post.author_id = author.id
         LEFT JOIN PostLike as like_table ON
            like_table.user_id = (SELECT id FROM UserProfile WHERE email = $1)
        AND like_table.post_id = post.id

         LEFT JOIN groups as community on post.group_id = community.id
         left join groupsubscriber g on community.id = g.group_id

         left join userprofile owner on post.owner_id = owner.id
         left join friendrequests f on owner.id = f.subscribed

         join userprofile sub on sub.id = f.subscriber or  g.user_id = sub.id
where sub.email = $1 
and post.creation_date > $3
order by creation_date
limit $2
	`
)

var (
	GetCommentByIdQuery = `
		SELECT id,
			   post_id,
			   text_content,
			   creation_date,
			   change_date,
			   is_deleted
		FROM Comment as c
		WHERE c.id = $1 AND is_deleted = false;
	`

	GetCommentSenderInfoQuery = `
		SELECT first_name, last_name, link, url as avatar_url
 				FROM UserProfile as p
 						 LEFT JOIN Photo as ph ON avatar_id = ph.id
 				JOIN comment AS c ON c.user_id = p.id AND c.id = $1
	`

	GetReplyReceiverInfoQuery = `
		SELECT first_name, last_name, link, url as avatar_url
 				FROM UserProfile as p
 						 LEFT JOIN Photo as ph ON avatar_id = ph.id
 				JOIN comment AS c ON c.reply_to = p.id AND c.id = $1
	`

	GetCommentsByPostIdQuery = `
		SELECT id,
			   post_id,
			   text_content,
			   creation_date,
			   change_date,
			   is_deleted,
			   CASE WHEN (SELECT email FROM UserProfile WHERE id = user_id) = $4 THEN true else false end as is_author
		FROM Comment as c
		WHERE c.post_id = $1 AND is_deleted = false AND creation_date > $3
		ORDER BY creation_date
		LIMIT $2
	`

	CreateCommentQuery = `
		INSERT INTO Comment(post_id, user_id, reply_to, text_content, creation_date, change_date, is_deleted)
		VALUES ($1, (SELECT id FROM UserProfile WHERE email = $2), (SELECT id FROM UserProfile WHERE link = $3), $4, $5, $5, FALSE) RETURNING id
	`

	DeleteCommentQuery = `
		UPDATE Comment
		SET is_deleted = true
		WHERE id = $1;	
	`

	UpdateCommentQuery = `
		UPDATE Comment
		SET text_content = $2, change_date = $3
		WHERE id = $1;
	`

	HasNextCommentsQuery = `
		SELECT CASE WHEN COUNT(*) <> 0 THEN true ELSE false END FROM Comment WHERE creation_date > $1 and is_deleted = false and post_id = $2
	`

	IsCommentAuthor = `
		SELECT CASE WHEN (SELECT COUNT(*) FROM comment AS c JOIN UserProfile as u on u.id = c.user_id WHERE c.id = $1 AND email = $2) = 0 then False else True end;
	`

	IsCommentDeleted = `
		SELECT is_deleted FROM Comment WHERE id = $1
	`
)
