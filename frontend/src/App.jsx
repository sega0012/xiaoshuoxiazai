import {React,useState} from 'react';
import { Spin,Button,  Form, Input ,ConfigProvider} from 'antd';
import {Greet,GetDetail} from "../wailsjs/go/main/App";




function App() {
	const [form] = Form.useForm();
	
	const [resultText, setResultText] = useState("下载状态");
	const [loading, setLoading] = useState(false);
	const [disable, setDisable] = useState(false);
	const onFinish = async (values) => {
  console.log('Success:', values);
	setResultText("小说下载中，请耐心等待")
	setDisable(true)
	setLoading(true)
	const msg = await Greet(values)
	console.log(msg)
	setResultText(msg)
	setDisable(false)
	setLoading(false)
	
};
const getdetail= async (e)=>{
	if(e.target.value==""){
		return
	}
	
	console.log(e.target.value)
	const xiaoshuo=await GetDetail(e.target.value)
	console.log(xiaoshuo);
	form.setFieldsValue({ title: xiaoshuo.Title,headUrl:xiaoshuo.HeadUrl,content:xiaoshuo.Content,nextPage:xiaoshuo.NextPage,breakFlag:xiaoshuo.BreakFlag });
	
}
const onFinishFailed = (errorInfo) => {
  console.log('Failed:', errorInfo);
	
	
};
	return (
    <ConfigProvider
  theme={{
	token: {
      colorTextDisabled:"white",
    },
    components: {
      Form: {
        labelColor:"white",
      },
      
    },
  }}
><br />
<h1 style={{color:'#3498db'}}>粘贴小说第一章网址</h1>
  <Form form={form}
    name="basic"
    labelCol={{
      span: 10,
    }}
    wrapperCol={{
      span: 20,
    }}
    style={{
      maxWidth: 780,
    }}
    initialValues={{
      remember: true,
    }}
    onFinish={onFinish}
    onFinishFailed={onFinishFailed}
    autoComplete="off"
  >
    <Form.Item
      label="文件名"
      name="name"
      rules={[
        {
          required: true,
          message: '请输入文件名',
        },
      ]}
    >
      <Input />
    </Form.Item>
    <Form.Item
      label="第一章地址"
      name="firstUrl"
      rules={[
        {
          required: true,
          message: '请输入第一章地址',
        },
      ]}
    >
      <Input onChange={(e)=>getdetail(e)}/>
    </Form.Item>

     
    <Form.Item
      label="网站前缀"
      name="headUrl"
      rules={[
        {
          required: true,
          message: '请输入网站前缀',
        },
      ]}
    >
      <Input />
    </Form.Item>
     <Form.Item
      label="小说本体DOM，使用JQ语法，比如#content"
      name="content"
      rules={[
        {
          required: true,
          message: '小说内容DOM',
        },
      ]}
    >
      <Input />
    </Form.Item>
     <Form.Item
      label="章节标题DOM,比如.detail_info h1"
      name="title"
      rules={[
        {
          required: true,
          message: '请输入章节标题Xpath',
        },
      ]}
    >
      <Input />
    </Form.Item>
     <Form.Item
      label="下一页DOM 比如.textinfo span:nth-child(4) a"
      name="nextPage"
      rules={[
        {
          required: true,
          message: '下一页URL的Xpath',
        },
      ]}
    >
      <Input />
    </Form.Item>
     <Form.Item
      label="终止标记"
      name="breakFlag"
      rules={[
        {
          required: true,
          message: '终止标记',
        },
      ]}
    >
      <Input />
    </Form.Item>
    
    
    
    <Form.Item
      wrapperCol={{
        offset: 8,
        span: 16,
      }}
    >
      <Button type="primary" htmlType="submit" disabled={disable}>
        点击下载
      </Button>
    <Spin size="large" spinning={loading} />
    </Form.Item>
  </Form>




<div id="result" className="result">{resultText}</div>
</ConfigProvider>
)
};
export default App;