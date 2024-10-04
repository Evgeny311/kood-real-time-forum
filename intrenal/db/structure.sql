CREATE TABLE
  IF NOT EXISTS user (
    id text PRIMARY KEY UNIQUE NOT NULL,
    username text UNIQUE NOT NULL,
    email text UNIQUE NOT NULL,
    password text NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    date_of_birth text NOT NULL,
    first_name text,
    last_name text,
    gender text,
    image_url text
  );

CREATE TABLE
  IF NOT EXISTS reaction (
    id text PRIMARY KEY UNIQUE NOT NULL,
    emoji text NOT NULL,
    description text NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS category (
    id text PRIMARY KEY UNIQUE NOT NULL,
    name text UNIQUE,
    description text
  );

CREATE TABLE
  IF NOT EXISTS parentReaction (
    id text PRIMARY KEY UNIQUE NOT NULL,
    user_id text NOT NULL,
    reaction_id text NOT NULL,
    parent_id text NOT NULL,
    parent_type text NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (reaction_id) REFERENCES reaction (id)
  );

-- POST DB INITIALIZATION --
CREATE TABLE
  IF NOT EXISTS post (
    id text PRIMARY KEY UNIQUE NOT NULL,
    title text NOT NULL,
    content text NOT NULL,
    user_id text NOT NULL,
    image_url text NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id)
  );

CREATE TABLE
  IF NOT EXISTS postCategory (
    id text PRIMARY KEY UNIQUE NOT NULL,
    post_id text NOT NULL,
    category_id text NOT NULL,
    FOREIGN KEY (post_id) REFERENCES post (id),
    FOREIGN KEY (category_id) REFERENCES category (id)
  );

CREATE TABLE
  IF NOT EXISTS comment (
    id text PRIMARY KEY UNIQUE NOT NULL,
    user_id text NOT NULL,
    post_id text NOT NULL,
    parent_id text,
    content text NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (post_id) REFERENCES post (id),
    FOREIGN KEY (parent_id) REFERENCES comment (id)
  );

-- CHAT DB INITIALIZATION --
CREATE TABLE
  IF NOT EXISTS chat (id text PRIMARY KEY UNIQUE NOT NULL);

CREATE TABLE
  IF NOT EXISTS chatUser (
    id text PRIMARY KEY UNIQUE NOT NULL,
    user_id text NOT NULL,
    chat_id text NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user (id),
    FOREIGN KEY (chat_id) REFERENCES chat (id)
  );

CREATE TABLE
  IF NOT EXISTS textLines (
    id text PRIMARY KEY UNIQUE NOT NULL,
    chat_id text NOT NULL,
    user_id text NOT NULL,
    content text NOT NULL,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES chat (id),
    FOREIGN KEY (user_id) REFERENCES user (id)
  );

-- {{{
-- INSERT INTO
--   user (
--     id,
--     username,
--     email,
--     password,
--     date_of_birth,
--     first_name,
--     last_name,
--     gender,
--     image_url
--   )
-- VALUES
--   (
--     'user_1',
--     'john_doe',
--     'john.doe@example.com',
--     'hashed_password_1',
--     '2024-01-13',
--     'John',
--     'Doe',
--     'Male',
--     'http://example.com/images/john_doe.png'
--   ),
--   (
--     'user_2',
--     'jane_smith',
--     'jane.smith@example.com',
--     'hashed_password_2',
--     '2024-01-13',
--     'Jane',
--     'Smith',
--     'Female',
--     'http://example.com/images/jane_smith.png'
--   ),
--   (
--     'user_3',
--     'alex_jones',
--     'alex.jones@example.com',
--     'hashed_password_3',
--     '2024-01-13',
--     'Alex',
--     'Jones',
--     'Non-Binary',
--     'http://example.com/images/alex_jones.png'
--   ) ON CONFLICT DO NOTHING;
-- INSERT INTO
--   category (id, name, description)
-- VALUES
--   (
--     'category_1',
--     'Technology',
--     'Discussions related to technology and gadgets.'
--   ),
--   (
--     'category_2',
--     'Cooking',
--     'Share and discover new recipes.'
--   ),
--   (
--     'category_3',
--     'Travel',
--     'Share travel stories and advice.'
--   ) ON CONFLICT DO NOTHING;
-- }}}
INSERT INTO
  reaction (id, emoji, description)
VALUES
  ('reaction_1', '😀', 'Happy'),
  ('reaction_2', '😢', 'Sad'),
  ('reaction_3', '👍', 'Like'),
  ('reaction_4', '❤️', 'Love'),
  ('reaction_5', '😂', 'Laugh'),
  ('reaction_6', '😡', 'Angry'),
  ('reaction_7', '😎', 'Cool'),
  ('reaction_8', '😴', 'Sleepy'),
  ('reaction_9', '😇', 'Angel'),
  ('reaction_10', '🤔', 'Thinking'),
  ('reaction_11', '😷', 'Sick'),
  (
    'reaction_12',
    '🤣',
    'Rolling on the floor laughing'
  ),
  ('reaction_13', '🙄', 'Rolling eyes'),
  ('reaction_14', '😮', 'Surprised'),
  ('reaction_15', '🤐', 'Silent'),
  ('reaction_16', '😋', 'Yummy'),
  ('reaction_17', '😪', 'Sleepy again'),
  ('reaction_18', '😜', 'Winking'),
  ('reaction_19', '😳', 'Blushing'),
  ('reaction_20', '🤩', 'Excited'),
  ('reaction_21', '😵', 'Dizzy'),
  ('reaction_22', '🤗', 'Hugging'),
  ('reaction_23', '🤫', 'Shushing'),
  ('reaction_24', '🙃', 'Upside down'),
  ('reaction_25', '😋', 'Delicious'),
  ('reaction_26', '🙏', 'Praying'),
  ('reaction_27', '🎉', 'Celebrating'),
  ('reaction_28', '🎈', 'Balloons'),
  ('reaction_29', '🌟', 'Starstruck'),
  ('reaction_30', '🔥', 'Fire'),
  ('reaction_31', '🌈', 'Rainbow'),
  ('reaction_32', '🍕', 'Pizza'),
  ('reaction_33', '🍦', 'Ice Cream'),
  ('reaction_34', '🍔', 'Burger'),
  ('reaction_35', '🍟', 'French Fries'),
  ('reaction_36', '🍣', 'Sushi'),
  ('reaction_37', '🍩', 'Donut'),
  ('reaction_38', '🍰', 'Cake'),
  ('reaction_39', '☕', 'Coffee'),
  ('reaction_40', '🍺', 'Beer'),
  ('reaction_41', '🍷', 'Wine'),
  ('reaction_42', '🍹', 'Cocktail'),
  ('reaction_43', '🍸', 'Martini'),
  ('reaction_44', '🥂', 'Cheers'),
  ('reaction_45', '🍾', 'Champagne'),
  ('reaction_46', '🥤', 'Soft Drink'),
  ('reaction_47', '🍼', 'Milk'),
  ('reaction_48', '🥛', 'Glass of Milk'),
  ('reaction_49', '🧀', 'Cheese'),
  ('reaction_50', '🍖', 'Meat'),
  ('reaction_51', '🥓', 'Bacon'),
  ('reaction_52', '🍗', 'Chicken'),
  ('reaction_53', '🍤', 'Shrimp'),
  ('reaction_54', '🍳', 'Egg'),
  ('reaction_55', '🍔', 'Hamburger'),
  ('reaction_56', '🍟', 'French Fries'),
  ('reaction_57', '🍕', 'Pizza'),
  ('reaction_58', '🥪', 'Sandwich'),
  ('reaction_59', '🌭', 'Hot Dog'),
  ('reaction_60', '🍝', 'Spaghetti'),
  ('reaction_61', '🍜', 'Ramen'),
  ('reaction_62', '🍲', 'Stew'),
  ('reaction_63', '🍛', 'Curry'),
  ('reaction_64', '🍙', 'Onigiri'),
  ('reaction_65', '🍚', 'Rice'),
  ('reaction_66', '🍘', 'Rice Cake'),
  ('reaction_67', '🍢', 'Oden'),
  ('reaction_68', '🍡', 'Dango'),
  ('reaction_69', '🍧', 'Shaved Ice'),
  ('reaction_70', '🍨', 'Ice Cream'),
  ('reaction_71', '🍦', 'Soft Serve Ice Cream'),
  ('reaction_72', '🍰', 'Cake'),
  ('reaction_73', '🎂', 'Birthday Cake'),
  ('reaction_74', '🧁', 'Cupcake'),
  ('reaction_75', '🥧', 'Pie'),
  ('reaction_76', '🍫', 'Chocolate Bar'),
  ('reaction_77', '🍬', 'Candy'),
  ('reaction_78', '🍭', 'Lollipop'),
  ('reaction_79', '🍯', 'Honey Pot'),
  ('reaction_80', '🍪', 'Cookie'),
  ('reaction_81', '🥠', 'Fortune Cookie'),
  ('reaction_82', '🍩', 'Doughnut'),
  ('reaction_83', '🍺', 'Beer Mug'),
  ('reaction_84', '🍻', 'Clinking Beer Mugs'),
  ('reaction_85', '🥂', 'Clinking Glasses'),
  ('reaction_86', '🥃', 'Tumbler Glass'),
  ('reaction_87', '🥤', 'Cup with Straw'),
  ('reaction_88', '🍹', 'Tropical Drink'),
  ('reaction_89', '🍸', 'Martini Glass'),
  ('reaction_90', '🥢', 'Chopsticks'),
  ('reaction_91', '🍽️', 'Fork and Knife with Plate'),
  ('reaction_92', '🍴', 'Fork and Knife'),
  ('reaction_93', '🥄', 'Spoon'),
  ('reaction_94', '🔪', 'Kitchen Knife'),
  ('reaction_95', '🏺', 'Amphora'),
  ('reaction_96', '🌿', 'Herb'),
  ('reaction_97', '🍀', 'Four Leaf Clover'),
  ('reaction_98', '🍁', 'Maple Leaf'),
  ('reaction_99', '🍂', 'Fallen Leaf'),
  ('reaction_100', '🍃', 'Leaf Fluttering in Wind') ON CONFLICT DO NOTHING;