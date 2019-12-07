use sms;

insert ignore into recv_log (mdn,receive_mdn,external_id,message,created)
                      values(4153007931,5102301414,'SM21e093678bfb460c8540d84c9f0ed5a2','Hi there!',utc_timestamp());