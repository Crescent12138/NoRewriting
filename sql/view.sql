create view user_pass_problem
    (user_id,sum_problem,rating_1000,rating_2000,rating_3000,over_3000)
    as select user_id ,count(*),
              count(if(rating <= 1000,1,null)),
              count(if(rating <= 2000 && rating >1000,1,null)),
              count(if(rating <= 3000 && rating >2000,1,null)),
              count(if(rating > 3000,11,null))
    from submission group by user_id
