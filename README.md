**Search Algorithm By Keyword In Golang**

1.Like Operator

    SELECT ID,question_name,category_id,topic_id,subject_id,subcategory_id FROM esdb.question  where question_name like '%vector%';

    +----+-----------------+--------+--------+--------+--------+
    93 rows in set (0.0188 sec)

2.Regex


    SELECT ID,question_name,category_id,topic_id,subject_id,subcategory_id FROM esdb.question  where question_name REGEXP 'vector';

    +----+-----------------+--------+--------+--------+--------+
    93 rows in set (0.0177 sec)

3.Priority

    Alter table question ADD FULLTEXT(question_name);

    SELECT ID,question_name,category_id,topic_id,subject_id,subcategory_id FROM esdb.question  where MATCH(question_name) AGAINST('vector');

    +----+-----------------+--------+--------+--------+--------+
    19 rows in set (0.0841 sec)
4.Id Priority

    SELECT ID,question_name,category_id,topic_id,subject_id,subcategory_id FROM esdb.question  where MATCH(question_name) AGAINST('vector') and (category_id=1 OR topic_id=1 OR subject_id=1 OR subcategory_id=1);

    +----+-----------------+--------+--------+--------+--------+
    19 rows in set (0.0015 sec)



    SELECT ID,question_name,category_id,topic_id,subject_id,subcategory_id FROM esdb.question  where MATCH(question_name) AGAINST('vector') and (category_id=0 OR topic_id=1 OR subject_id=0 OR subcategory_id=1);

    +----+-----------------+--------+--------+--------+--------+
    14 rows in set (0.0013 sec)
