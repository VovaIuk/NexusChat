-- Insert 4 users
INSERT INTO Users (tag, username, password) VALUES
('alice_dev', 'Alice', '$2a$10$tU/g4T.CtPoujDyyJJpIR.kDAbqq569iHCg.XlR9XdYVn921h8okG'),  -- alice1234
('bob_codes', 'Bob', '$2a$10$5bQ6YPZpIX34mzywXSQo6OSWQL8XQyVx26RuzCEayqQq3GRuI4PJa'),    -- bob12345
('charlie_js', 'Charlie', '$2a$10$2TP0vHoJGaxiAQk6DlAIdu.A.JOBq3djs3tJGed0dKr5SvNeovL3K'), -- char1234
('diana_py', 'Diana', '$2a$10$OEDTqgRG9lpuoR0IAq9DE.fXXSrxj/8dE..5W6/P4aj7ASOMSsGka');   -- diana1234

-- Insert 4 private chats (each chat ID will be auto-generated)
INSERT INTO Chats DEFAULT VALUES; -- Chat 1: Alice & Bob
INSERT INTO Chats DEFAULT VALUES; -- Chat 2: Alice & Charlie
INSERT INTO Chats DEFAULT VALUES; -- Chat 3: Alice & Diana
INSERT INTO Chats DEFAULT VALUES; -- Chat 4: Bob & Charlie

-- Link users to chats (User_Chat relationships)
INSERT INTO User_Chat (user_id, chat_id) VALUES
(1, 1), (2, 1),  -- Chat 1: Alice + Bob
(1, 2), (3, 2),  -- Chat 2: Alice + Charlie
(1, 3), (4, 3),  -- Chat 3: Alice + Diana
(2, 4), (3, 4);  -- Chat 4: Bob + Charlie

-- Insert messages with realistic timestamps (UTC)
INSERT INTO Messages (user_id, chat_id, text, time) VALUES
-- Chat 1 (Alice & Bob): 3 messages
(1, 1, 'Hey Bob, are we still on for the meeting today?', '2025-12-10 09:15:00+00'),
(2, 1, 'Yes! I''ll share the agenda by noon.', '2025-12-10 09:17:30+00'),
(1, 1, 'Perfect, thanks!', '2025-12-10 09:18:45+00'),

-- Chat 2 (Alice & Charlie): 2 messages
(1, 2, 'Did you review the PR I sent?', '2025-12-11 14:30:00+00'),
(3, 2, 'Just left comments. The logic looks solid!', '2025-12-11 14:45:20+00'),

-- Chat 3 (Alice & Diana): 1 message
(4, 3, 'Alice, can you share the design docs?', '2025-12-12 11:05:10+00'),

-- Chat 4 (Bob & Charlie): 4 messages
(2, 4, 'Found a bug in the auth module. Can you help?', '2025-12-13 16:20:00+00'),
(3, 4, 'Sure, send me the stack trace.', '2025-12-13 16:22:15+00'),
(2, 4, 'Here it is: [ERROR] Token validation failed', '2025-12-13 16:25:40+00'),
(3, 4, 'Ah, I see the issue. Fixing it now!', '2025-12-13 16:30:55+00');