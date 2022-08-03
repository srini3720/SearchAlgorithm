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

5.Multiple Tables Search  (Join)

    SELECT  
      question.question_name,
      question.ID,
      question.subcategory_id,
      question.category_id,
      question.subject_id,
      question.topic_id,
      subjects.subject_name,
      topic.topic_name,
      subcategory.subcategory_name,
      category.category_name
    FROM question 
      LEFT JOIN subcategory ON question.subcategory_id = subcategory.ID 
      LEFT JOIN category ON question.category_id = category.ID 
      LEFT JOIN topic ON question.topic_id = topic.ID 
      LEFT JOIN subjects ON question.subcategory_id = subjects.ID 
    WHERE 
      MATCH(question.question_name) AGAINST('vector')
      OR (MATCH(subcategory.subcategory_name) AGAINST('vector')   OR subcategory.ID =1 )
      OR (MATCH(category.category_name) AGAINST('vector')   OR category.ID =1)
      OR (MATCH(subjects.subject_name) AGAINST('vector')   OR subjects.ID = 1)
      OR (MATCH(topic.topic_name) AGAINST('vector')  OR topic.ID =1);



 
    +----+-----------------+--------+--------+--------+--------+
       morethat 10 lakhs rows in set (73 sec)
    
5.Multiple Tables Search  (Priority)

                (SELECT '1' as priority,ID,question_name as suggestion FROM question WHERE MATCH(question_name) AGAINST('vector'))
      UNION ALL (SELECT '2' as priority,ID,subject_name as suggestion FROM subjects WHERE  (MATCH(subject_name AGAINST('vector')) )

      UNION ALL (SELECT '3' as priority,ID,topic_name as suggestion FROM topic WHERE   (MATCH(topic_name) AGAINST('vector')) )

      UNION ALL (SELECT '4' as priority,ID,category_name as suggestion FROM category WHERE  (MATCH(category_name) AGAINST('vector')) )

      UNION ALL (SELECT '5' as priority,ID,subcategory_name as suggestion FROM subcategory WHERE   (MATCH(subcategory_name) AGAINST('vector')) )

      ORDER BY priority 

    +----+-----------------+--------+--------+--------+--------+
       111 rows in set (0.05 sec)