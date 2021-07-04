SELECT username, SUM(deleted_count) AS deleted_count
    FROM vmd_statistics GROUP BY username ORDER BY deleted_count DESC;
