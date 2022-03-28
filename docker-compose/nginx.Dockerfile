# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM nginx
COPY html /usr/share/nginx/html

COPY server.conf /etc/nginx/conf.d/default.conf
