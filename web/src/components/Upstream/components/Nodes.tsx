import { MinusCircleOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Col, Form, Input, InputNumber, Row } from 'antd';
import React from 'react';
import { useIntl } from 'umi';

import { removeBtnStyle } from '..';

type Props = {
  readonly?: boolean;
};

const Component: React.FC<Props> = ({ readonly }) => {
  const { formatMessage } = useIntl();

  return (
    <Form.List
      name="submitNodes"
      initialValue={[{ host: undefined, port: undefined, weight: undefined }]}
    >
      {(fields, { add, remove }) => (
        <>
          <Form.Item
            label={formatMessage({ id: 'page.upstream.form.item-label.node.domain.or.ip' })}
            style={{ marginBottom: 0 }}
          >
            {fields.map((field, index) => (
              <Row style={{ marginBottom: 10 }} gutter={16} key={index}>
                <Col xs={9} sm={12} md={9} lg={10} xl={7} xxl={6}>
                  <Form.Item
                    label={formatMessage({ id: 'page.upstream.step.host' })}
                    style={{ marginBottom: 0 }}
                    name={[field.name, 'host']}
                    rules={[
                      {
                        required: true,
                        message: formatMessage({
                          id: 'page.upstream.step.input.domain.name.or.ip',
                        }),
                      },
                      {
                        // eslint-disable-next-line no-useless-escape
                        pattern: new RegExp(/^\*?[0-9a-zA-Z-._\[\]:]+$/),
                        message: formatMessage({
                          id: 'page.route.form.itemRulesPatternMessage.domain',
                        }),
                      },
                    ]}
                  >
                    <Input
                      placeholder={formatMessage({ id: 'page.upstream.step.domain.name.or.ip' })}
                      disabled={readonly}
                    />
                  </Form.Item>
                </Col>
                <Col md={5} lg={5} xl={5} xxl={4}>
                  <Form.Item
                    style={{ marginBottom: 0 }}
                    name={[field.name, 'port']}
                    label={formatMessage({ id: 'page.upstream.step.port' })}
                  >
                    <InputNumber
                      placeholder={formatMessage({ id: 'page.upstream.step.port' })}
                      disabled={readonly}
                      min={1}
                      max={65535}
                    />
                  </Form.Item>
                </Col>
                <Col md={5} lg={5} xl={5} xxl={4}>
                  <Form.Item
                    style={{ marginBottom: 0 }}
                    name={[field.name, 'weight']}
                    label={formatMessage({ id: 'page.upstream.step.weight' })}
                    rules={[
                      {
                        required: true,
                        message: formatMessage({ id: 'page.upstream.step.input.weight' }),
                      },
                    ]}
                    initialValue={1}
                  >
                    <InputNumber
                      placeholder={formatMessage({ id: 'page.upstream.step.weight' })}
                      disabled={readonly}
                      min={0}
                      max={1000}
                    />
                  </Form.Item>
                </Col>
                <Col style={{ ...removeBtnStyle }}>
                  {!readonly && (
                    <MinusCircleOutlined
                      data-cy={`upstream-node-minus-${index}`}
                      onClick={() => remove(field.name)}
                    />
                  )}
                </Col>
              </Row>
            ))}
          </Form.Item>
          {!readonly && (
            <Form.Item wrapperCol={{ offset: 3 }}>
              <Button type="dashed" onClick={add} data-cy="add-node">
                <PlusOutlined />
                {formatMessage({ id: 'component.global.add' })}
              </Button>
            </Form.Item>
          )}
        </>
      )}
    </Form.List>
  );
};

export default Component;
