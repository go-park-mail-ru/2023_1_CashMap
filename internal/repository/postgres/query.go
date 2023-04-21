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
	CreateGroup       = ``
	UpdateGroupLink   = ``
	UpdateGroupAvatar = ``
	DeleteGroup       = ``
	GroupSubscribe    = ``
	GroupUnsubscribe  = ``
	AddManager        = ``
)
