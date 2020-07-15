import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../providers/product.dart';
import '../providers/products.dart';

class EditProductPage extends StatefulWidget {
  static const String routeName = '/edit-product';

  @override
  _EditProductPageState createState() => _EditProductPageState();
}

class _EditProductPageState extends State<EditProductPage> {
  final _priceFocusNode = FocusNode();
  final _descriptionFocusNode = FocusNode();
  final _imageUrlFocusNode = FocusNode();
  final _imageUrlController = TextEditingController();
  final _form = GlobalKey<FormState>();
  Product _editedProduct =
      Product(id: null, title: '', price: 0, description: '', imageUrl: '');
  bool _isInit = true;
  bool _isLoading = false;

  @override
  void initState() {
    super.initState();
    _imageUrlFocusNode.addListener(_updateImageUrl);
  }

  @override
  void didChangeDependencies() {
    super.didChangeDependencies();
    if (_isInit) {
      _isInit = false;
      final int productId = ModalRoute.of(context).settings.arguments as int;
      if (productId == null) return;
      _editedProduct =
          Provider.of<Products>(context, listen: false).findByID(productId);
      _imageUrlController.text = _editedProduct.imageUrl;
    }
  }

  void _updateImageUrl() {
    if (!_imageUrlFocusNode.hasFocus) {
      setState(() {});
    }
  }

  Future<void> _saveForm() async {
    final bool isValid = _form.currentState.validate();
    if (!isValid) return;
    _form.currentState.save();
    setState(() {
      _isLoading = true;
    });
    final Products products = Provider.of<Products>(context, listen: false);
    try {
      if (_editedProduct.id != null) {
        await products.updateByID(_editedProduct.id, _editedProduct);
        Navigator.of(context).pop();
      } else {
        await products.add(_editedProduct);
        Navigator.of(context).pop();
      }
    } catch (error) {
      await showDialog<Null>(
        context: context,
        builder: (context) => AlertDialog(
          title: Text('An error occured!'),
          content: Text('Something went wrong.'),
          actions: <Widget>[
            FlatButton(
              child: Text('Ok'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            )
          ],
        ),
      );
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  void dispose() {
    super.dispose();
    _priceFocusNode.dispose();
    _descriptionFocusNode.dispose();
    _imageUrlFocusNode.removeListener(_updateImageUrl);
    _imageUrlFocusNode.dispose();
    _imageUrlController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Edit product'),
        actions: <Widget>[
          IconButton(
            icon: Icon(Icons.save),
            onPressed: _saveForm,
          ),
        ],
      ),
      body: _isLoading
          ? Center(
              child: CircularProgressIndicator(),
            )
          : Padding(
              padding: const EdgeInsets.all(15),
              child: Form(
                key: _form,
                child: ListView(
                  children: <Widget>[
                    TextFormField(
                        initialValue: _editedProduct.title,
                        decoration: InputDecoration(labelText: 'Title'),
                        textInputAction: TextInputAction.next,
                        onFieldSubmitted: (_) {
                          FocusScope.of(context).requestFocus(_priceFocusNode);
                        },
                        validator: (value) {
                          if (value.isEmpty) return 'Please enter a value.';
                          return null;
                        },
                        onSaved: (title) {
                          _editedProduct = _editedProduct.clone(title: title);
                        }),
                    TextFormField(
                      initialValue: _editedProduct.price.toString(),
                      decoration: InputDecoration(labelText: 'Price'),
                      keyboardType: TextInputType.number,
                      textInputAction: TextInputAction.next,
                      focusNode: _priceFocusNode,
                      onFieldSubmitted: (_) {
                        FocusScope.of(context)
                            .requestFocus(_descriptionFocusNode);
                      },
                      validator: (value) {
                        if (value.isEmpty) return 'Please enter a value.';
                        if (double.tryParse(value) == null)
                          return 'Please enter a valid number.';
                        if (double.parse(value) < 0)
                          return 'Please enter a number greater than zero.';
                        return null;
                      },
                      onSaved: (price) {
                        _editedProduct =
                            _editedProduct.clone(price: double.parse(price));
                      },
                    ),
                    TextFormField(
                      initialValue: _editedProduct.description,
                      decoration: InputDecoration(labelText: 'Description'),
                      maxLines: 3,
                      keyboardType: TextInputType.multiline,
                      focusNode: _descriptionFocusNode,
                      validator: (value) {
                        if (value.isEmpty) return 'Please enter a description.';
                        if (value.length < 10)
                          return 'The description should be at least 10 characters long.';
                        return null;
                      },
                      onSaved: (description) {
                        _editedProduct =
                            _editedProduct.clone(description: description);
                      },
                    ),
                    Row(
                      crossAxisAlignment: CrossAxisAlignment.end,
                      children: <Widget>[
                        Container(
                          width: 100,
                          height: 100,
                          margin: EdgeInsets.only(
                            top: 8,
                            right: 10,
                          ),
                          decoration: BoxDecoration(
                              border: Border.all(
                                width: 1,
                                color: Colors.grey,
                              ),
                              borderRadius: BorderRadius.circular(15)),
                          child: _imageUrlController.text.isEmpty
                              ? Center(child: const Text('Enter a URL'))
                              : ClipRRect(
                                  borderRadius: BorderRadius.circular(15),
                                  child: FittedBox(
                                    child:
                                        Image.network(_imageUrlController.text),
                                    fit: BoxFit.cover,
                                  ),
                                ),
                        ),
                        Expanded(
                          child: TextFormField(
                            decoration: InputDecoration(labelText: 'Image URL'),
                            keyboardType: TextInputType.url,
                            textInputAction: TextInputAction.done,
                            controller: _imageUrlController,
                            focusNode: _imageUrlFocusNode,
                            onFieldSubmitted: (_) => _saveForm(),
                            validator: (value) {
                              if (value.isEmpty)
                                return 'Please enter an image URL.';
                              var urlPattern =
                                  r"(https?|ftp)://([-A-Z0-9.]+)(/[-A-Z0-9+&@#/%=~_|!:,.;]*)?(\?[A-Z0-9+&@#/%=~_|!:‌​,.;]*)?";
                              bool isValidUrl =
                                  new RegExp(urlPattern, caseSensitive: false)
                                      .hasMatch(value);
                              if (!isValidUrl)
                                return 'Please enter a valid URL.';
                              return null;
                            },
                            onSaved: (imageUrl) {
                              _editedProduct =
                                  _editedProduct.clone(imageUrl: imageUrl);
                            },
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ),
    );
  }
}
