package render

const fullProductMatch = `select distinct on (p.name) p.* from products as p left join product_descriptions as pds on p.name = pds.product_name 
  where name ~* '?.*$' OR type ~* '?.*$' OR pds.value ~* '?.*$'`
