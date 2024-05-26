import React, { useState } from 'react';
import axios from 'axios';

const InvoiceForm = () => {
  const [amount, setAmount] = useState('');
  const [department, setDepartment] = useState('');
  const [managerApproval, setManagerApproval] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const invoiceData = {
      amount: parseFloat(amount),
      department,
      manager_approval: managerApproval,
    };

    try {
      const response = await axios.post('http://localhost:8080/invoice', invoiceData);
      console.log('Invoice submitted:', response.data);
    } catch (error) {
      console.error('There was an error submitting the invoice!', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Amount:</label>
        <input
          type="number"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Department:</label>
        <input
          type="text"
          value={department}
          onChange={(e) => setDepartment(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Manager Approval:</label>
        <input
          type="checkbox"
          checked={managerApproval}
          onChange={(e) => setManagerApproval(e.target.checked)}
        />
      </div>
      <button type="submit">Submit Invoice</button>
    </form>
  );
};

export default InvoiceForm;
