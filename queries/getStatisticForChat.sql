SELECT username, deleted_count
FROM vmd_statistics WHERE chat_id=? ORDER BY deleted_count DESC;