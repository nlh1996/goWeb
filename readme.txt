1.�޸�demo/database/conn.go��mongo�����ӵ�ַ��Session,err = mgo.Dial("mongodb://admin:admin@115.159.77.155:11000?maxPoolSize=100")����Ĭ�������Ʒ�������һ���������ݿ��ַ����auth��֤
2.��docker�����ļ��зŵ��������ϣ�����docker�����ļ�Ŀ¼�� ִ��"docker build -t goweb ."����Զ�����docker���񣨸þ�����go���������б���demo��û��ʹ��Ԥ�ȱ����demoֱ���������ж������ļ��ķ����� ��
3.������� ִ�С�docker run -d --name=goweb -p 8080:8000 goweb����������������������8080�˿�ӳ�䵽�����ڲ���8000�˿ڣ�8080�˿ڿ����Զ���
4.web����Ķ˿���8000��������demo/router/router.go�������޸ģ�Ȼ����dockerfile�п����޸ĺ�Ķ˿ڡ�
5.���� �ڷ������ն�ʹ�� curl http://localhost:8080/v1/hello/aa ���м򵥲���
6.��������qq:1217286494 ��ϵ�ҡ�