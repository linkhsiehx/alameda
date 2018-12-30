import {QueryResponse} from 'src/logs/api/v2'
import {SearchStatus} from 'src/types/logs'

export const fluxResponse: QueryResponse = {
  tables: [
    {
      id: 'd7cc1e08-4b17-4309-885b-6798402bdae2',
      data: [
        [
          '',
          'result',
          'table',
          'appname',
          'facility',
          'host',
          'severity',
          'message',
          'procid',
          'timestamp',
        ],
        [
          '',
          '',
          '0',
          'testbot-2',
          'NTP subsystem',
          'user.local',
          'alert',
          'this is the first log message\\n',
          '92273',
          '1539285296911000000',
        ],
        [
          '',
          '',
          '0',
          'testbot-2',
          'cron',
          'user.local',
          'warning',
          'this is the second\\n',
          '92273',
          '1539285306912000000',
        ],
        [
          '',
          '',
          '0',
          'testbot-2',
          'cron',
          'user.local',
          'err',
          'this is the third message\\n',
          '92273',
          '1539285306912000000',
        ],
        [
          '',
          '',
          '0',
          'testbot-2',
          'lpr',
          'user.local',
          'crit',
          'This is the fourth message\\n',
          '92273',
          '1539285316917000000',
        ],
        [
          '',
          '',
          '0',
          'testbot-2',
          'lpr',
          'user.local',
          'err',
          'This is the fifth message\\n',
          '92273',
          '1539285316917000000',
        ],
      ],
      name: '',
      result: '',
      groupKey: {},
      dataTypes: {
        '': '#datatype',
        result: 'string',
        table: 'long',
        appname: 'string',
        facility: 'string',
        host: 'string',
        severity: 'string',
        message: 'string',
        procid: 'string',
        timestamp: 'long',
      },
    },
  ],
  status: SearchStatus.Loaded,
}
