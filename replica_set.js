cfg = {
    "_id": "RS",
    "members": [{
        "_id": 0,
        "host": "rs1:27041"
    },
        {
            "_id": 1,
            "host": "rs2:27042"
        },
        {
            "_id": 2,
            "host": "rs3:27043"
        }
    ]
};
rs.initiate(cfg);