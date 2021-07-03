UPDATE vmd_statistics
SET deleted_count = deleted_count + 1,
    modified_at=?
WHERE chat_id=? AND user_id=?;
