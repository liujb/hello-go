<?php

$token = "%7B%22ticket%22%3A%22CkAn_wSQtzyM2fB5-wXfH9t0enQye5Tz1qy_VoQ0vHFMzdGqgzAMxvF3-a5z0faoZ_RlRrBhCtVKGxlDfPdlG4PlJoF_4Hdg3-eECBDeR3ADoZYsiI7AN9u-7wiL6FRSQzywcWv3Uu3Zd70Pl-E_uJMwVmGVq86L_BYLkrMJ_s99x7Bx4nWV_JETK-tje5nnMwAA__8%3D%22%7D";

$json = json_encode(array("ticket" => $token));

echo urlencode($json)."\r";