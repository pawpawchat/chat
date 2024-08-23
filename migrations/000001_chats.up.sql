

CREATE TABLE chats (
    chat_id BIGSERIAL PRIMARY KEY,
    title VARCHAR(32) NOT NULL,
    number_of_members INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE messages (
    message_id BIGSERIAL PRIMARY KEY,
    body VARCHAR(1024) NOT NULL,
    chat_id BIGINT NOT NULL,
    sent_at TIMESTAMP NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT false,
    FOREIGN KEY (chat_id) REFERENCES chats(chat_id)
);


CREATE TABLE chat_members (
    member_id BIGINT NOT NULL,
    username VARCHAR(16) NOT NULL,
    chat_id BIGINT NOT NULL,
    role VARCHAR(16) NOT NULL,
    PRIMARY KEY (member_id, chat_id),
    FOREIGN KEY (chat_id) REFERENCES chats(chat_id) ON DELETE CASCADE
);


-- Функция для триггера new_chat_member
-- Срабатывает при добавлении/исключения пользователя из чата
CREATE OR REPLACE FUNCTION increace_number_chat_members()
RETURNS TRIGGER AS $$
BEGIN

    -- При добавлении увеличивает число участников в таблице chats
    IF TG_OP = 'INSERT' THEN
        UPDATE chats SET number_of_members = number_of_members + 1 WHERE chat_id = NEW.chat_id;
        RETURN NEW;
    -- При удалении уменьшает число участников в таблице chats
    ELSEIF TG_OP = 'DELETE' THEN
        UPDATE chats SET number_of_members = number_of_members - 1 WHERE chat_id = OLD.chat_id;
        RETURN OLD;
    END IF;
    RETURN NULL;

END;
$$ LANGUAGE plpgsql;


-- Триггер для таблицы chats, следит за количеством участников
CREATE TRIGGER new_chat_member
AFTER INSERT OR DELETE ON chat_members
FOR EACH ROW
EXECUTE FUNCTION increace_number_chat_members();
